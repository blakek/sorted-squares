CC ?= clang
SOURCES = sorted-squares.c
OBJECTS = $(SOURCES:.c=.o)
TARGET = sorted-squares
CFLAGS = -framework Foundation

default: $(TARGET)
all: sorted-squares

dev: | clean all
	./$(TARGET)

$(TARGET):
	$(CC) $(CFLAGS) -o sorted-squares sorted-squares.c
	@echo "Successfully built"

clean:
	$(RM) *.o
	$(RM) $(TARGET)

.c.o:
	$(CC) $(CFLAGS) -c $< -o $@
