try
	get application id "com.stairways.keyboardmaestro.engine"
on error err_msg number err_num
	return "Keyboard Maestro not found. First install it and then use this workflow."
end try

tell application id "com.stairways.keyboardmaestro.engine"
	gethotkeys with asstring and getall
end tell
