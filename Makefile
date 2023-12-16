TARGET = commands-gui
TARGET_DIR = build
TARGET_PATH = ./$(TARGET_DIR)/$(TARGET)

app: clean prepare
	go build -o $(TARGET_PATH) main.go

run:
	$(TARGET_PATH)

clean:
	rm -f $(TARGET_PATH)

prepare:
	if [[ ! -d ./$(TARGET_DIR) ]]; then mkdir ./$(TARGET_DIR); fi
