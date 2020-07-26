if test -f ".env"; then
    source ./.env
fi

FOLDER_NAME=_utmist.gitlab.io

cd onedeath
lua main.lua $ONEDRIVE_FOLDER_LINK
find . -type f -name '*.exf' -delete
rm -rf ../content
mv $FOLDER_NAME/* ../
rm -rf $FOLDER_NAME cookies.txt
cd ..
