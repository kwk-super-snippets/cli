#!/usr/bin/env bash

set -ef -o pipefail
KWK_VERSION=v1.0.17
BUILD_NUMBER=$1

echo "**** kwk-cli build ${KWK_VERSION}+${BUILD_NUMBER} *****"

ARCH=amd64

# TESTING
go test ./app
go test ./ui/dlg

# PREP OUTPUT
releasePath=/builds/${KWK_VERSION}
tmp=/builds/temp
binPath=${releasePath}/bin
npmPath=${releasePath}/npm

if [[ -d "${tmp}" ]]; then
    rm -fr ${tmp}
else
    mkdir ${tmp}
fi

if [[ ! -d "${releasePath}" ]]; then
    mkdir ${releasePath}
fi

if [[ ! -d "${binPath}" ]]; then
    mkdir ${binPath}
fi

if [[ ! -d "${npmPath}" ]]; then
    mkdir ${npmPath}
fi

# COMPILING
function compile(){
  os=$1
  file="kwk-${os}-${ARCH}"

  # COMPILE
  binary=${tmp}/bin/${file}
  env GOOS=${os} GOARCH=${ARCH} go build -ldflags "-X main.version=${KWK_VERSION} -X main.build=${BUILD_NUMBER}" -x -o ${binary}

  # ZIP
  zipped=${binPath}/${file}.tar.gz
  tar cvzf ${zipped} -C ${tmp}/bin ${file}

  # CHECKSUM
  sha1sum ${zipped} > ${zipped}.sha1
}

sed -i -- "s/RELEASE_VERSION/${KWK_VERSION}/" ./main.go
compile linux
compile darwin
compile windows

# CREATE NPM
npmTemp=${tmp}/npm
rm -fr ${npmTemp}
mkdir ${npmTemp}

cp -R dist/npm/. ${npmTemp}
cp -R ${tmp}/bin/. ${npmTemp}/bin
sed -i -- "s/RELEASE_VERSION/${KWK_VERSION}/" ${npmTemp}/package.json
tree ${npmTemp}

# CREATE NPM TAR
echo "CREATING NPM ARCHIVE"
npmTar=${npmPath}/kwk-cli-npm.tar.gz
tar cvzf ${npmTar} -C ${npmTemp} .
sha1sum ${npmTar} > ${npmTar}.sha1

# CLEAN-UP
rm -fr ${tmp}

# TODO: WARNING REMOVE THIS WHEN OPEN SOURCING
export AWS_ACCESS_KEY_ID=AKIAJRJBQNMZWLG653WA
export AWS_SECRET_ACCESS_KEY=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
export DEFAULT_REGION="us-east-1"

aws s3 cp /builds/${KWK_VERSION} s3://kwk-cli/${KWK_VERSION} --recursive --acl public-read