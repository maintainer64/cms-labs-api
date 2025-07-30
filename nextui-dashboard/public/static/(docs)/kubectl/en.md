# Connecting to Kubernetes Clusters

This guide will help you connect to Kubernetes clusters using two main methods: via the topology web interface or using
the `kubectl` utility on your local computer. We'll break down each step in detail to make the process simple and
straightforward. Before you begin, ensure you have access to the lab environment and the necessary credentials.

## 1. Connecting via the Topology Interface

This method allows you to interact with the cluster directly in your browser without installing additional software.
It's ideal for quick access and requires no setup on your device.

### Connection Steps:

1. **Select a lab assignment**: Navigate to the lab assignments section and choose the desired one.
2. **Wait for system login**: After selecting the assignment, the system will automatically log you in. This may take a
   few seconds.
3. **Wait for the topology to load**: On the `/topology/` page, wait until all topology components are fully loaded.
   You'll see a visual diagram with available elements.

### Monitoring the Loading Status:

If you want to track the loading progress:

1. Click the menu button in the top-left corner of the `/topology/` page.
2. Select **Open Topology Logs**. This will open a log where you can view the progress and any potential errors.

### Interacting with Components:

- Select the desired component (e.g., a cluster node) in the topology.
- A web terminal will automatically open for interacting with the equipment. You can now execute commands directly in
  your browser.

This method is convenient for beginners, but for advanced work, using `kubectl` is recommended.

## 2. Connecting Using the kubectl Utility

`kubectl` is a powerful command-line tool for managing Kubernetes clusters. To use it, you need to install the utility
on your computer and configure authentication using the Krew plugin. We'll go through the process step by step for
different operating systems.

### 2.1. Installing kubectl

First, download and install `kubectl`. This is the core utility, and further steps are impossible without it.

#### Installation Steps:

1. **Download the binary file**:
    - Go to the official Kubernetes website: [releases.kubernetes.io](https://kubernetes.io/releases/).
    - Select the version suitable for your OS (Linux, macOS, or Windows). A stable version is recommended (e.g.,
      v1.28.x).

2. **Installation on Linux or macOS**:
    - Make the file executable:
      ```bash
      chmod +x ./kubectl
      ```
    - Move the file to a directory in your PATH environment variable (e.g., `/usr/local/bin/`):
      ```bash
      sudo mv ./kubectl /usr/local/bin/kubectl
      ```

3. **Installation on Windows**:
    - Download the `kubectl.exe` file.
    - Place it in a convenient folder, such as `C:\Program Files\kubectl`.
    - Add this folder to your PATH environment variable:
        - Open "Settings" > "System" > "About" > "Advanced system settings" > "Environment Variables".
        - Under "System variables", find PATH, edit it, and add the folder path.

4. **Verify the installation**:
    - Run the command:
      ```bash
      kubectl version --client
      ```
    - If the installation was successful, you'll see the `kubectl` version (e.g., Client Version: v1.28.0). If errors
      occur, check your PATH and permissions.

### 2.2. Installing Krew and Authentication Plugins

Krew is a plugin manager for `kubectl` that simplifies installing extensions. We'll use it for the OIDC authorization
plugin and a special plugin for accessing the training cluster. If Krew isn't installed yet, follow the steps below.

#### Installing Krew:

1. **For Linux/macOS**:
    - Install dependencies if needed: `git` and `curl`.
    - Run the installation script:
      ```bash
      (
        set -x; cd "$(mktemp -d)" &&
        OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
        ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
        KREW="krew-${OS}_${ARCH}" &&
        curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz" &&
        tar zxvf "${KREW}.tar.gz" &&
        ./"${KREW}" install krew
      )
      ```
    - Add Krew to your PATH: add the line `export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"` to `~/.bashrc` or
      `~/.zshrc` and restart the terminal.

2. **For Windows**:
    - Install using PowerShell or follow the official
      guide: [krew.sigs.k8s.io](https://krew.sigs.k8s.io/docs/user-guide/setup/install/).

3. **Verify Krew installation**:
    - Run:
      ```bash
      kubectl krew version
      ```
    - If the version is displayed, Krew is ready to use.

#### Installing Plugins:

1. **Install the OIDC authorization plugin**:
    - Run:
      ```bash
      kubectl krew install oidc-login
      ```
    - This will enable authentication using your training account.

2. **Install the cluster access plugin**:
    - Run:
      ```bash
      kubectl krew install --manifest-url https://gitlab.com/a10869/kubectl-cms/-/raw/main/krew.yaml
      ```

3. **Configuring contexts and access**:
    - Run:
      ```bash
      kubectl cmslab
      ```
    - This will add the necessary contexts for your training cluster.

4. **Verify the configuration**:
    - Check available contexts:
      ```bash
      kubectl config get-contexts
      ```
    - You'll see a list of contexts, including the training one (e.g., named `urfu`).

#### Example Cluster Interaction:

- To list pods in a specified namespace, use:
  ```bash
  kubectl get pods -n {{namespace}} --context urfu
  ```
