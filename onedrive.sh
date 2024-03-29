if test -f ".env"; then
    . ./.env
fi

FOLDER_NAME=$ONEDRIVE_FOLDER_NAME

cd onedeath
lua main.lua $ONEDRIVE_FOLDER_LINK
find . -type f -name '*.exf' -delete
rm -rf ../content* ../insertions* ../static/*.pdf ../static/images/profilepics/
mv $FOLDER_NAME/config.yaml ../
mv $FOLDER_NAME/content ../content_base
mv $FOLDER_NAME/insertions ../insertions_base
mv $FOLDER_NAME/static/*.pdf ../static/
mv $FOLDER_NAME/static/images/profilepics ../static/images/
rm -rf $FOLDER_NAME cookies.txt
cd ..
