# prePressFileChecker
Проверяет ресурсы (линкованные картинки) перед автоматической сборкой печатного прайс-листа.

Берет все строки из excel начинающиеся на `<Pict>`
пример: `<Pict>D:\WORK\TRADE LOCK\Прайс_2023\Фото в прайс\45907_01`

Извлекает имя файла 45907_01, добавляет расширение и проверяет файл на существование в указанной папке.
Путь собирается из: image_folder_path / 45907_01 . image_extension

## Конфигурация
Для указания пути к файлу конфигурации запускать с флагом -p
-p string
        Path for YAML configuration file (default "./config.yaml")
        
### В конфигурации задается:

путь к файлу с данными
excel_file_path: C:\dev\excel\Прайс_ТЛ_FINAL_LIZA.xlsx

путь к папке с фотографиями для вставки в прайс
image_folder_path: C:\dev\image

расширение файлов с фотографиями
image_extension: jpg
