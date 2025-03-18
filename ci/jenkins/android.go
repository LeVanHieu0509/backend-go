package jenkins

/*
1. Restrict where this project can be run
*/
/*
#Build IOS --> ANDROID(current)
cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/custom-config/
export FLUTTER_VERSION_NAME=$(yq eval '.versionDevName' version-store.yaml)
export FLUTTER_BUILD_NUMBER=$(yq eval '.dev-currentBuildNumber' version-store.yaml)
export FLUTTER_VERSION_FULL="${FLUTTER_VERSION_NAME}+${FLUTTER_BUILD_NUMBER}"

cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/fwd-cpr-app/

git reset --hard
git clean -fd .
git clean -fdX
git checkout release/dev
git fetch
git pull

yq eval -i '.version = env(FLUTTER_VERSION_FULL)' pubspec.yaml

cd lib/
cp /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/custom-config/env.dart.sit env.dart
cd ..

# Install project dependencies
flutter clean
flutter doctor -v
flutter pub get

# Build .ipa with Obfucscate option
flutter build apk --obfuscate --split-debug-info=/Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr/deployment/logs/fwd-cpr-app

# Upload to Google Driver
cd ~/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/fwd-cpr-app/build/app/outputs/flutter-apk
mv app-release.apk BIT-DEV-FWD-CPR_"$FLUTTER_VERSION_NAME"_"$FLUTTER_BUILD_NUMBER".apk
gdrive files upload --parent 1Yp_LqaYDJdL2TZjuJdNPJNtWBlb55flD BIT-DEV-FWD-CPR_"$FLUTTER_VERSION_NAME"_"$FLUTTER_BUILD_NUMBER".apk

# Clear git workspace
cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/fwd-cpr-app/
git reset --hard
git clean -fd .
git clean -fdX

curl -s -X POST https://api.telegram.org/bot811113517:AAGfW9c5p3NkMKa09dCjQWFUt8ce8cncQdc/sendMessage -d chat_id=-4554825141 -d text="BIT-FWD-CPR (DEV): Đã build xong bản $FLUTTER_VERSION_NAME ($FLUTTER_BUILD_NUMBER)"
*/
