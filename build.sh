#!/usr/bin/env bash
set -e

# Convert Windows paths to Unix-style (MSYS2/Git Bash handles this)
OPENCV_DIR="C:\\opencv\\build"
OPENCV_INCLUDE="$OPENCV_DIR\\install\\include"
OPENCV_LIB="$OPENCV_DIR\\install\\x64\\mingw\\staticlib"
# OPENCV_LIB="$OPENCV_DIR\\install\\x64\\mingw\\lib"

# Ensure Go is in PATH
if ! command -v go &> /dev/null; then
  echo "Go not found in PATH"
  exit 1
fi

# Ensure MinGW-w64 GCC is in PATH
if ! command -v x86_64-w64-mingw32-gcc &> /dev/null; then
  echo "MinGW-w64 GCC not found in PATH"
  exit 1
fi

# Output binary name
OUTPUT="build/myapp.exe"

# Flags
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export CGO_ENABLED=1
export CGO_CPPFLAGS="-I$OPENCV_INCLUDE"
export CGO_LDFLAGS="-L$OPENCV_LIB -static \
  -lopencv_calib3d4120 \
  -lopencv_core4120 \
  -lopencv_dnn4120 \
  -lopencv_features2d4120 \
  -lopencv_highgui4120 \
  -lopencv_imgcodecs4120 \
  -lopencv_imgproc4120 \
  -lopencv_objdetect4120 \
  -lopencv_photo4120 \
  -lopencv_video4120 \
  -lopencv_videoio4120 \
  -lopencv_flann4120 \
  -lzlib \
  -lIlmImf \
  -llibwebp \
  -llibtiff \
  -llibjpeg-turbo \
  -llibopenjp2 \
  -llibpng \
  -llibprotobuf \
  -lcomdlg32 \
  -loleaut32 \
  -lole32 \
  -luuid \
  -luser32 \
  -lgdi32"

# Build
go build -tags customenv -ldflags="-extldflags '-static'" -o "$OUTPUT"
# go build -tags customenv -o "$OUTPUT"
# go build -tags customenv -ldflags="-extldflags=-static" -o "$OUTPUT"

echo "✅ Build complete: $OUTPUT"
