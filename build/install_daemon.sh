#!/bin/bash

SERVICE_NAME="dicdh"
DESCRIPTION="Dynamic IP Cloudflare DNS Handler"
SERVICE_PATH="/usr/bin/dicdh"

IS_ACTIVE=$(sudo systemctl is-active $SERVICE_NAME)
if [ "$IS_ACTIVE" == "active" ]; then
    echo "dicdh is already running, restarting the service."
    sudo systemctl restart $SERVICE_NAME
    echo "dicdh has been restarted."
else
    sudo cat > /etc/systemd/system/$SERVICE_NAME.service << EOF
[Unit]
Description=$DESCRIPTION
After=network.target
[Service]
ExecStart=$SERVICE_PATH
Restart=on-failure
[Install]
WantedBy=multi-user.target
EOF
    sudo systemctl daemon-reload 
    sudo systemctl enable $SERVICE_NAME
    sudo systemctl start $SERVICE_NAME
fi

echo "done."

exit 0