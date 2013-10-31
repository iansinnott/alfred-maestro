try
	get application id "com.stairways.keyboardmaestro.engine"
on error err_msg number err_num
	return "There was an error..."
end try

tell application id "com.stairways.keyboardmaestro.engine"
	gethotkeys with asstring
end tell
