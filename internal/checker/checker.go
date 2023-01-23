// Package checker проверяет ресурсы (линкованные картинки) перед автоматической сборкой печатного
// прайс-листа.
// Берет все строки из excel начинающиеся на <Pict>
// пример: <Pict>D:\WORK\TRADE LOCK\Прайс_2023\Фото в прайс\45907_01
// Извлекает имя файла 45907_01, добавляет расширение и проверяет файл на существование в указанной папке.
// Путь собирается из: image_folder_path / 45907_01 . image_extension
//
// В конфигурации (config.yaml) задается:
// путь к файлу с данными
// excel_file_path: C:\dev\excel\Прайс_ТЛ_FINAL_LIZA.xlsx
// путь к папке с фотографиями для вставки в прайс
// image_folder_path: C:\dev\image
// расширение файлов с фотографиями
// image_extension: jpg
package checker

import (
	"fmt"

	"github.com/atrian/prePressFileChecker/internal/config"
	"github.com/atrian/prePressFileChecker/internal/fileChecker"
	"github.com/atrian/prePressFileChecker/internal/xmlReader"
	"github.com/atrian/prePressFileChecker/internal/zipReader"
	"github.com/atrian/prePressFileChecker/pkg/logger"
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
	fmt.Println("easservice@gmail.com for any questions")
}
