Vagrant.configure("2") do |config|
  config.vm.box = "alvistack/ubuntu-24.04"
  
  config.vm.provider :libvirt do |libvirt|
    libvirt.driver = "kvm"
    libvirt.memory = 8192
    libvirt.cpus = 4
    libvirt.graphics_type = "none"
    libvirt.memorybacking :access, :mode => "shared"
  end

  config.vm.synced_folder ".", "/workspace", type: "virtiofs"

  # Enable native X11 forwarding
  config.ssh.forward_x11 = true

  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    
    # Install Go, X11 auth, and required Electron GUI libraries
    apt-get install -y golang-go git curl apt-transport-https xauth dbus-x11 libnss3 libxss1 libasound2t64 libatk-bridge2.0-0 libgtk-3-0 libgbm1 libsecret-1-0

    # Install VSCodium
    wget -qO - https://gitlab.com/paulcarroty/vscodium-deb-rpm-repo/raw/master/pub.gpg | gpg --dearmor -o /etc/apt/keyrings/vscodium-archive-keyring.gpg
    echo 'deb [ signed-by=/etc/apt/keyrings/vscodium-archive-keyring.gpg ] https://download.vscodium.com/debs vscodium main' > /etc/apt/sources.list.d/vscodium.list
    apt-get update && apt-get install -y codium

    # Pre-install extensions
    su - vagrant -c "codium --install-extension golang.go --install-extension k--kato.intellij-idea-keybindings --install-extension Google.geminicodeassist"
  SHELL
end
