#!/bin/bash
nohup ./venv/bin/python3 main.py >> nohup.out & export PID=$!
	&& echo $PID > run.pid \
	&& echo "#!/bin/bash" > kill.sh \
	&& echo "kill -9 $PID" >> "kill.sh" \
	&& chmod +x kill.sh

