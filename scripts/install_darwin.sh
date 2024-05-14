#!/bin/bash

while getopts ":v:p" opt; do
  case $opt in
  p)
    DESTINATION="$OPTARG"
    ;;
  v)
    VERSION="$OPTARG"
    ;;
  \?)
    echo "Invalid option -$OPTARG" >&2
    ;;
  esac
done

if [ -z "$DESTINATION" ]; then
  DESTINATION="/usr/local/bin"
fi

if [ -z "$VERSION" ]; then
  # Get latest version from github
  VERSION=$(curl -s https://api.github.com/repos/Parallels/prl-devops-service/releases/latest | grep -o '"tag_name": "[^"]*"' | cut -d ' ' -f 2 | tr -d '"')
fi

if [[ ! $VERSION == release-* ]]; then
  VERSION="release-$VERSION"
fi

ARCHITECTURE=$(uname -m)
if [ "$ARCHITECTURE" = "aarch64" ]; then
  ARCHITECTURE="arm64"
fi

OS=$(uname -s)
OS=$(echo "$OS" | tr '[:upper:]' '[:lower:]')

DOWNLOAD_URL="https://github.com/Parallels/prl-devops-service/releases/download/$VERSION/prldevops--$OS-$ARCHITECTURE.tar.gz"

echo "Downloading prldevops $VERSION from $DOWNLOAD_URL"
curl -sL $DOWNLOAD_URL -o prldevops.tar.gz

echo "Extracting prldevops"
tar -xzf prldevops.tar.gz

if [ ! -d "$DESTINATION" ]; then
  echo "Creating destination directory: $DESTINATION"
  mkdir -p "$DESTINATION"
fi

if [ -f "$DESTINATION/prldevops" ]; then
  echo "Removing existing prldevops"
  sudo rm "$DESTINATION/prldevops"
fi
echo "Moving prldevops to $DESTINATION"
sudo mv prldevops $DESTINATION/prldevops
sudo chmod +x $DESTINATION/prldevops

if [ "$OS" = "darwin" ]; then
  sudo xattr -d com.apple.quarantine $DESTINATION/prldevops
fi


echo "Cleaning up"
rm prldevops.tar.gz