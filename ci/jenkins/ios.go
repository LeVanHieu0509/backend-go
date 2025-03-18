package jenkins

/*
# Build IOS(current) --> ANDROID
curl -s -X POST https://api.telegram.org/bot811113517:AAGfW9c5p3NkMKa09dCjQWFUt8ce8cncQdc/sendMessage -d chat_id=-4554825141 -d text="BIT-FWD-CPR (DEV): Bắt đầu build"

cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/fwd-cpr-app

flutter clean
git reset --hard
git clean -fd .
git clean -fdX
git checkout release/dev
git fetch
git pull

# Defined session env
cd ../custom-config/
export FLUTTER_VERSION_NAME=$(yq eval '.versionDevName' version-store.yaml)
export FLUTTER_BUILD_NUMBER=$(yq eval '.dev-nextBuildNumber' version-store.yaml)
export FLUTTER_VERSION_FULL="${FLUTTER_VERSION_NAME}+${FLUTTER_BUILD_NUMBER}"
export DEV_IOS_PROVISIONING=$(yq eval '.PROVISIONING_PROFILE_SPECIFIER' cert-dev-ios-store.yaml)
export DEV_ASC_PUBLIC_ID=$(yq eval '.PUBLIC_ID' cert-dev-ios-store.yaml)
export DEV_PRODUCT_BUNDLE_IDENTIFIER=$(yq eval '.PRODUCT_BUNDLE_IDENTIFIER' cert-dev-ios-store.yaml)
export DEV_APPLE_ID=$(yq eval '.APPLE_ID' cert-dev-ios-store.yaml)
#!<key>CFBundleName</key>
export DEV_IOS_CFBUNDLE_NAME=$(yq eval '.CFBUNDLE_NAME' cert-dev-ios-store.yaml)
export DEV_IOS_ARCHIVE_PATH=/Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/fwd-cpr-app/ios/_archive/simp.xcarchive
export DEV_IOS_OPTIONS_PLIST_PATH=/Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/custom-config/DEVOptionsPlist.plist
export DEV_IOS_IPA_PATH=/Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/fwd-cpr-app/ios/_ipa/$DEV_IOS_CFBUNDLE_NAME.ipa
export LOCAL_U=$(yq eval '.u' .local.authorization.yaml)
export LOCAL_P=$(yq eval '.p' .local.authorization.yaml)


# Update project version
cd ../fwd-cpr-app
yq eval -i '.version = env(FLUTTER_VERSION_FULL)' pubspec.yaml
cd lib/
cp /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/custom-config/env.dart.sit env.dart
cd ..

# Setup pre-actions
flutter clean
flutter pub get
cd ios && pod install
cp ../../custom-config/env.prepare.dev.example .env.prepare
sh prepare-env.sh

# Stage 'build'
xcodebuild -sdk iphoneos -configuration Release -workspace Runner.xcworkspace -scheme Runner build | xcpretty -s -r junit --no-color

# Stage 'archive'
xcodebuild -workspace Runner.xcworkspace \
		   -scheme Runner \
           clean archive \
           -sdk iphoneos \
           -configuration Release \
           -archivePath ${DEV_IOS_ARCHIVE_PATH} \
           | xcpretty -s -r junit --no-color

xcodebuild -exportArchive \
		   -archivePath ${DEV_IOS_ARCHIVE_PATH} \
  		   -exportOptionsPlist ${DEV_IOS_OPTIONS_PLIST_PATH} \
           -exportPath ${DEV_IOS_IPA_PATH} \
           | xcpretty -s -r junit --no-color

# Stage 'release'
xcrun altool --validate-app \
			 -f "${DEV_IOS_IPA_PATH}/${DEV_IOS_CFBUNDLE_NAME}".ipa \
             -t ios \
             -u "${LOCAL_U}" -p "${LOCAL_P}"

xcrun altool --upload-package "${DEV_IOS_IPA_PATH}/${DEV_IOS_CFBUNDLE_NAME}".ipa \
			 -t ios \
             -u "${LOCAL_U}" -p "${LOCAL_P}" \
             --bundle-short-version-string "${FLUTTER_VERSION_NAME}" \
             --bundle-version "${FLUTTER_BUILD_NUMBER}" \
             --asc-public-id "${DEV_ASC_PUBLIC_ID}" \
             --bundle-id "${DEV_PRODUCT_BUNDLE_IDENTIFIER}" \
             --apple-id "${DEV_APPLE_ID}"

### Clean workspace
# Create new version
cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/custom-config/
yq eval -i '.dev-nextBuildNumber = (env(FLUTTER_BUILD_NUMBER) | tonumber + 1)' version-store.yaml
yq eval -i '.dev-currentBuildNumber = (env(FLUTTER_BUILD_NUMBER))' version-store.yaml


# Clear git workspace
cd /Users/nhancao/Bitelle/Jenkins-Workdir/fwd-cpr-app/deployment/sit/fwd-cpr-app/
git reset --hard
git clean -fd .
git clean -fdX
*/
