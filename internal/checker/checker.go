package checker

import (
	"fmt"
	"prePressFilesChecker/internal/config"
	"prePressFilesChecker/internal/fileChecker"
	"prePressFilesChecker/internal/xmlReader"
	"prePressFilesChecker/internal/zipReader"
	"prePressFilesChecker/pkg/logger"
)

type App struct {
	config *config.AppConfig
	logger logger.Logger
}

func New() *App {
	appLog := logger.NewZapLogger()

	appConf, err := config.NewConfig(appLog)
	if err != nil {
		appLog.Fatal("Can't load app configuration", err)
	}

	return &App{
		config: appConf,
		logger: appLog,
	}
}

func (a *App) Run() {

	// открываем excel файл как zip архив
	zip := zipReader.New()

	data, err := zip.UnzipFile(a.config.ExcelFilePath)
	if err != nil {
		a.logger.Error("Can't read data from excel file", err)
		defer data.Close()
	}

	// разбираем данные из файла в структуру
	xmlData := xmlReader.Sst{}
	err = xmlData.ParseXML(data)
	if err != nil {
		a.logger.Error("Can't parse XML data", err)
	}

	pictures := xmlData.GetPics()

	// закрываем ресурсы архива
	zip.CloseArchive()

	fch := fileChecker.New(a.config.ImageExtension, a.config.ImageFolderPath)
	fch.Load(pictures).CheckFiles()

	fmt.Println("Pics NOT found:")

	for _, pict := range fch.FilesNotExist {
		fmt.Println(pict)
	}

	fmt.Println("==========================================")

	fmt.Println("Pics found in EXCEL:", len(pictures))
	fmt.Println("Pics found in DIR:", len(fch.FilesExist))
	fmt.Println("Pics NOT found in DIR:", len(fch.FilesNotExist))

}
