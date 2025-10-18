FROM fedora:42

# Set environment variables
ARG USERNAME=devuser
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Install packages and create non-root user
RUN dnf update -y && \
    dnf install -y \
        git \
        golang \
        cascadia-mono-nf-fonts \
        ripgrep \
        make \
        gcc \
        nvim \
        sudo \
        shadow-utils && \
    dnf clean all && \
    \
    # Create group and user
    groupadd --gid $USER_GID $USERNAME && \
    useradd --uid $USER_UID --gid $USER_GID -m $USERNAME && \
    echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# Switch to non-root user
USER $USERNAME

# Set the working directory
WORKDIR /home/$USERNAME

# Make sure to pass your current user's UID/GID in .devcontainer.json
# {
#  "build": {
#    "dockerfile": "Dockerfile",
#    "args": {
#      "USER_UID": "1000",
#      "USER_GID": "1000"
#    }
#  },
#  "remoteUser": "devuser"
#}

