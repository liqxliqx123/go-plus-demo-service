#!/bin/bash
#
# Generate test coverage statistics for Go packages.
#

set -e

output() {
  color="32"
  if [[ "$2" -gt 0 ]]; then
    color="31"
  fi
  printf "\033[${color}m"
  echo $1
  printf "\033[0m"
}

workdir=".cover"
cover_mode="set"
kernel_name=$(uname -s)
packages=$(go list ./... | grep -v vendor)

show_help() {
cat << EOF
Generate test coverage statistics for Go packages.
  -- Command Flag --
  -h | --help                    Display this help and exit
  -m | --mode                    Set coverage mode. default is "set" (set|count|atomic)
  -d | --dir                     Set store coverage folder (default is ".cover")
  -- Command Action --
  tool                           Install go dependency tools like gocov or golint.
  testing                        Run go testing for all packages
  coverage                       Generate coverage xxx for all packages
  junit                          Generate coverage xml xxx for junit plugin
  lint                           Generate Lint xxx for all packages
  vet                            Generate Vet xxx for all packages
  cloc                           Generate Count Lines of Code xxx for all files
  all                            Execute coverage、junit、lint、vet and cloc xxx
Contribute and source at https://github.com/appleboy/golang-testing
EOF
exit 0
}

# add this function on your .bashrc file to echoed Go version
go_version() {
  if [ "$1" = "" ]; then
    match=1
  else
    match=2
  fi
  version=$(go version)
  regex="go(([0-9].[0-9]).[0-9])"
  if [[ $version =~ $regex ]]; then
    echo ${BASH_REMATCH[${match}]}
  fi
}

set_workdir() {
  workdir=$1
  test -d $workdir || mkdir -p $workdir
  coverage_xxx="$workdir/coverage.txt"
  coverage_xml_xxx="$workdir/coverage.xml"
  junit_xxx="$workdir/junit.txt"
  junit_xml_xxx="$workdir/xxx.xml"
  lint_xxx="$workdir/lint.txt"
  vet_xxx="$workdir/vet.txt"
  cloc_xxx="$workdir/cloc.xml"
}

install_dependency_tool() {
  goversion=$(go_version "gloabl")
  [ -d "${GOPATH}/bin" ] || mkdir -p ${GOPATH}/bin
  go get -u github.com/jstemmer/go-junit-xxx
  go get -u github.com/axw/gocov/gocov
  go get -u github.com/AlekSi/gocov-xml
  if [[ "$goversion" < "1.6" ]]; then
    output "Golint requires Go 1.6 or later."
  else
    go get -u github.com/golang/lint/golint
  fi
  curl https://raw.githubusercontent.com/AlDanial/cloc/master/cloc -o ${GOPATH}/bin/cloc
  chmod 755 ${GOPATH}/bin/cloc
}

errorNumber() {
  if [ "$1" -ne 0 ]; then
    error=$1
  fi
}

testing() {
  error=0
  output "Begin running testing."
  test -f ${junit_xxx} && rm -f ${junit_xxx}
  output "Running ${cover_mode} mode for coverage."

  for pkg in $packages; do
    compilefile="$workdir/compile_$(echo $pkg | tr / -)"
    f="$workdir/$(echo $pkg | tr / -).cover"
    output "Testing coverage xxx for ${pkg}"

    go test -c -coverpkg=./... -covermode=${cover_mode} -o ${compilefile} $pkg

    if [ -f "$compilefile" ]
    then
      go tool test2json -t ${compilefile} -test.v -test.coverprofile ${f}
    else
      echo "$compilefile not found."
    fi
    # go test -p 1 -v -cover -coverprofile=${f} -covermode=${cover_mode} $pkg | tee -a ${junit_xxx}
    # ref: http://stackoverflow.com/questions/1221833/bash-pipe-output-and-capture-exit-status
    errorNumber ${PIPESTATUS[0]}
  done

  output "Convert all packages coverage xxx to $coverage_xxx"
  echo "mode: $cover_mode" > "$coverage_xxx"
  grep -h -v "^mode:" "$workdir"/*.cover >> "$coverage_xxx"
  if [ "$error" -ne 0 ]; then
    output "Get Tesing Error Number Code: ${error}" ${error}
  fi
}

generate_cover_xxx() {
  gocov convert ${coverage_xxx} | gocov-xml > ${coverage_xml_xxx}
}

generate_junit_xxx() {
  cat ${junit_xxx} | go-junit-xxx > ${junit_xml_xxx}
}

generate_lint_xxx() {
  for pkg in $packages; do
    output "Go Lint xxx for ${pkg}"
    golint ${pkg} | tee -a ${lint_xxx}
  done

  # fix path error
  root_path=${PWD//\//\\/}
  [ "$kernel_name" == "Darwin" ] && sed -e "s/${root_path}\(\/\)*//g" -i '' ${lint_xxx}
  [ "$kernel_name" == "Linux" ] && sed -e "s/${root_path}\(\/\)*//g" -i ${lint_xxx}
}

generate_vet_xxx() {
  for pkg in $packages; do
    output "Go Vet xxx for ${pkg}"
    go vet -n -x ${pkg} | tee -a ${vet_xxx}
  done
}

generate_cloc_xxx() {
  cloc --by-file --xml --out=${cloc_xxx} --exclude-dir=vendor,Godeps,.cover .
}

# set default folder.
set_workdir $workdir

# Process command line...

[ $# -gt 0 ] || show_help

while [ $# -gt 0 ]; do
  case $1 in
    --help | -h)
      show_help
    ;;
    --mode | -m)
      shift
      cover_mode=$1
      test -z ${cover_mode} && show_help
      shift
      test -z $1 && show_help
      ;;
    --dir | -d)
      shift
      workdir=$1
      test -z ${workdir} && show_help
      set_workdir ${workdir}
      shift
      test -z $1 && show_help
      ;;
    tool)
      install_dependency_tool
      shift
      ;;
    testing)
      testing
      shift
      ;;
    coverage)
      generate_cover_xxx
      shift
      ;;
    junit)
      generate_junit_xxx
      shift
      ;;
    lint)
      generate_lint_xxx
      shift
      ;;
    vet)
      generate_vet_xxx
      shift
      ;;
    cloc)
      generate_cloc_xxx
      shift
      ;;
    all)
      testing
      generate_cover_xxx
      generate_junit_xxx
      generate_lint_xxx
      generate_vet_xxx
      generate_cloc_xxx
      shift
      ;;
    *)
      show_help ;;
  esac
done

if [[ "$error" -gt 0 ]]; then
  exit $error
fi
