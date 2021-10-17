# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.187.0/containers/go/.devcontainer/base.Dockerfile

ARG VARIANT="1.17"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN go install \
    github.com/volatiletech/sqlboiler/v4@latest \
    github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest \
    && go get -x \
    github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
    github.com/fatih/gomodifytags
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN chmod -R a+w /go/pkg

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1