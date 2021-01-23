##################################################
# OS: Ubuntu 20.04 LTS
##################################################

#!/bin/bash

##################################################
# Variables
##################################################
CURRENT=`date '+%Y%m%d-%H%M%S'`
LOG_DIR='/var/log/scripts'
LOG_FILE="setup-$(hostname)-${CURRENT}.log"
LOG_PATH="${LOG_DIR}/${LOG_FILE}"

##################################################
# Functions
##################################################
output_label() {
  echo '##################################################' >> $LOG_PATH
  echo "# $1" >> $LOG_PATH
  echo '##################################################' >> $LOG_PATH
}

output_stdout() {
  echo '--------------------------------------------------' >> $LOG_PATH
  echo "| [Command] $1" >> $LOG_PATH
  echo '--------------------------------------------------' >> $LOG_PATH
  eval $1 >> $LOG_PATH
}

##################################################
# Main
##################################################
### Create Log Directory
if [ ! -d "${LOG_DIR}" ]; then
  mkdir -p ${LOG_DIR}
fi

### Check User
if [ "${USER}"  != "root" ]; then
  echo 'You must be logged in as "root" to run this tool.'
  exit 1
fi

### OS Basic
output_label 'OS Basic'
output_stdout 'cat /etc/lsb-release'
output_stdout 'hostnamectl status'
output_stdout 'timedatectl status'
output_stdout 'localectl status'
output_stdout 'cat /etc/shells'
output_stdout 'cat /etc/sudoers'
output_stdout 'cat /etc/rsyslog.conf'

### Network
output_label 'Network'
output_stdout 'cat /etc/nsswitch.conf'
output_stdout 'cat /etc/snmp/snmpd.conf'
output_stdout 'ls -a /etc/netplan/*.yaml'

for f in `find /etc/netplan/*.yaml`; do
  output_stdout "cat ${f}"
done

for i in `netstat -ia | awk 'NR > 2 { print $1 }'`; do
  output_stdout "ethtool ${i}"
  output_stdout "ethtool -i ${i}"
done

### Disk
output_label 'Disk'
output_stdout 'df -Th'
output_stdout 'fdisk -l'
output_stdout 'pvs'
output_stdout 'vgs'
output_stdout 'lvs'

### Service
output_label 'Service'
output_stdout 'systemctl list-unit-files --type service'
output_stdout 'systemctl list-unit-files --type mount'
output_stdout 'systemctl list-unit-files --type slice'
output_stdout 'systemctl list-unit-files --type timer'
output_stdout 'systemctl list-unit-files --type socket'
output_stdout 'systemctl list-unit-files --type target'
output_stdout 'systemctl list-dependencies default.target'

### Kernel
output_label 'Kernel'
output_stdout 'sysctl --system'
output_stdout 'sysctl -a'
output_stdout 'cat /etc/sysctl.conf'
output_stdout 'ls -l /usr/lib/sysctl.d'

for f in `find /usr/lib/sysctl.d`; do
  output_stdout "cat ${f}"
done

### Package
output_label 'Package'
output_stdout 'dpkg -l'
output_stdout 'snap list'

### User Group
output_label 'User Group'
output_stdout 'cat /etc/passwd'
output_stdout 'cat /etc/group'

### NTP
output_label "NTP"
# output_stdout 'cat /etc/ntp.conf'
# output_stdout 'cat ntpq -p'
output_stdout 'cat /etc/chrony.conf'
output_stdout 'chronyc sources'
output_stdout 'chronyc tracking'

### DNS
output_label 'DNS'
output_stdout 'cat /etc/hosts'
output_stdout 'ls -l /etc/resolv.conf'
output_stdout 'cat /run/systemd/resolve/resolv.conf'
output_stdout 'cat /run/systemd/resolve/stub-resolv.conf'
output_stdout 'systemd-resolve --status'

### Security
output_label 'Security'
output_stdout 'ls -a /etc/pam.d'
output_stdout 'cat /etc/login.defs'
output_stdout 'cat /etc/hosts.allow'
output_stdout 'cat /etc/hosts.deny'

### Logrotate
output_label "Logrotate"
output_stdout 'cat /etc/logrotate.conf'
output_stdout 'ls -l /etc/logrotate.d'
for f in `find /etc/logrotate.d`; do
  output_stdout "cat ${f}"
done

### Job
output_label 'Job'
output_stdout 'cat /etc/crontab'
output_stdout 'cat /etc/cron.deny'
output_stdout 'cat /etc/at.deny'
output_stdout 'ls -l /etc/cron.daily'
output_stdout 'ls -l /etc/cron.weekly'
output_stdout 'ls -l /etc/cron.monthly'

### Hardware
output_label 'Hardware'

### Software
output_label 'Software'

exit 0
