# FrancisWiFi


## Prérequis

### Mise à jour et installation des prérequis

```bash
sudo apt update
sudo apt upgrade -y
sudo apt-get install build-essential hostapd iptables golang-go git dnsmasq -y
```

### Install create_ap

```bash
git clone https://github.com/oblique/create_ap.git
cd create_ap/
sudo make install
cd ..
```
### Install mes trucs
```bash
git clone https://github.com/mikelimazulu/franciswifi.git
cd franciswifi/
go mod init captive
go mod tidy
go build captive.go
sudo mv captive /usr/bin/
chmod +x update_iptables.sh
sudo mv update_iptables.sh /usr/bin
sudo mv staticContentFolder/ /etc

```

## Configuration des services

### iptables

`sudo nano /etc/systemd/system/iptables-config.service`

```
[Unit]
Description=Configure iptables for captive portal
Before=create-ap.service

[Service]
Type=oneshot
ExecStart=/usr/bin/update_iptables.sh
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target
```

### create_ap

`sudo nano /etc/systemd/system/create-ap.service`

```
[Unit]
Description=Create Wi-Fi Access Point
After=iptables-config.service
Before=captive-portal.service

[Service]
ExecStart=/usr/bin/create_ap -n wlan0 Guest_AC380
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target
```

### captive portal

`sudo nano /etc/systemd/system/captive-portal.service`

```
[Unit]
Description=Go Captive Portal Service
After=create-ap.service

[Service]
ExecStart=/usr/bin/captive
Restart=always

[Install]
WantedBy=multi-user.target
```

### enable services aux reboot

```
sudo systemctl enable iptables-config.service
sudo systemctl enable create-ap.service
sudo systemctl enable captive-portal.service
```
