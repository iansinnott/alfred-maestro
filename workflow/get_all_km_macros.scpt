try
	get application id "com.stairways.keyboardmaestro.engine"
on error err_msg number err_num
	return "Keyboard Maestro not found. You should install it first."
end try

tell application id "com.stairways.keyboardmaestro.engine"
	gethotkeys with asstring and getall
end tell
