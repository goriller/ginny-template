# /bin/bash
export LOG_PATH=logs/log.log
export LOG_LEVEL=info
export PRE_STOP="80s"

nohup ./app.bin -conf ../configs/config.yml > start.log 2>&1 &
