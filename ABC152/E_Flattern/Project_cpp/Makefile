CC=gcc
CC_FLAGS_DEBUG=-g -lm -O0 -o
CC_FLAGS_RELEASE=-fconstexpr-loop-limit=100000000 -lm -o

#PROJECT_NAME=E_Flattern
#SOURCE=$(PROJECT_NAME).cpp
PROJECT_NAME=main
SOURCE=main.cpp

ifeq ($(OS),Windows_NT)
	TARGET_FILE=$(PROJECT_NAME).exe
else
UNAME=${shell uname}
ifeq ($(UNAME),Darwin)
	TARGET_FILE=$(PROJECT_NAME)_for_mac
else
ifeq ($(UNAME),Linux)
	TARGET_FILE=$(PROJECT_NAME)_for_linux
else
	TARGET_FILE=$(PROJECT_NAME)
endif
endif
endif

all:	$(TARGET_FILE)

$(TARGET_FILE):	$(SOURCE)
	$(CC) $(CC_FLAGS_RELEASE)$(TARGET_FILE) $(SOURCE)

release:	$(SOURCE)
	$(CC) $(CC_FLAGS_RELEASE)$(TARGET_FILE) $(SOURCE)

debug:	$(SOURCE)
	$(CC) $(CC_FLAGS_DEBUG)debug_$(TARGET_FILE) $(SOURCE)

clean:
	rm $(TARGET_FILE)	# Can not run on Windows

.PHONY:	all clean
