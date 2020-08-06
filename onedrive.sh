if test -f ".env"; then
    source ./.env
fi

FOLDER_NAME=_utmist.gitlab.io

cd onedeath
lua main.lua $ONEDRIVE_FOLDER_LINK
find . -type f -name '*.exf' -delete
rm -rf ../content* ../templates*
mv $FOLDER_NAME/config.yaml ../
mv $FOLDER_NAME/content ../content_base
mv $FOLDER_NAME/templates ../templates_base
rm -rf $FOLDER_NAME cookies.txt
cd ..
