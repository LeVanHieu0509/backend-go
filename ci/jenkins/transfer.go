package jenkins

/*
cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/custom-config/
export FLUTTER_VERSION_NAME=$(yq eval '.versionName' version-store.yaml)
export FLUTTER_BUILD_NUMBER=$(yq eval '.uat-currentBuildNumber' version-store.yaml)
export FLUTTER_VERSION_FULL="${FLUTTER_VERSION_NAME}+${FLUTTER_BUILD_NUMBER}"

cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/transfer/bit/fwd-cpr-app
git reset --hard
git clean -fdX
git clean -fd .
flutter clean
git fetch origin
git checkout transfer/uat
git reset --hard
git clean -fd .
git status
git pull
echo "[#$FLUTTER_VERSION_FULL] $(git log -1 --pretty=format:%s)" > ../../commit_log

cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/transfer/fwd/fwd-cpr-app-azure
git reset --hard
git clean -fd .

git fetch origin
git checkout transfer/uat
git reset --hard
git clean -fdX
git clean -fd .
git fetch
git pull

cd ../../
rsync -av --progress --delete bit/fwd-cpr-app/ fwd/fwd-cpr-app-azure/ --exclude='.git*' --exclude='android/app/google-services.json' --exclude='ios/Runner/GoogleService-Info.plist'
cp bit/fwd-cpr-app/.gitignore fwd/fwd-cpr-app-azure/.gitignore

cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/transfer/fwd/fwd-cpr-app-azure
## Config version
yq eval -i '.version = env(FLUTTER_VERSION_FULL)' pubspec.yaml

git status
git add -A
git reset lib/env.dart
#git commit -F ../../commit_log
#git push origin transfer/uat

#flutter clean
#cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/transfer/bit/fwd-cpr-app
#flutter clean
curl -s -X POST https://api.telegram.org/bot811113517:AAGfW9c5p3NkMKa09dCjQWFUt8ce8cncQdc/sendMessage -d chat_id=-4554825141 -d text="BIT-FWD-CPR (TransferUAT): $FLUTTER_VERSION_NAME ($FLUTTER_BUILD_NUMBER) @cn7here please review!"
*/
