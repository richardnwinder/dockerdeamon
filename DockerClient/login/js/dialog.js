function formLogin_submit()
{
	var username = document.getElementById('usernamelog').value;
	var password = document.getElementById('passwordlog').value;
	var language = currentLanguage;
	console.log("username : " + username);
	console.log("password : " + password);
	console.log("language : " + language);
	// validate data here
	var validationError = "";
	if (username === "")
		validationError = "Username is blank ! Please enter username.\n";
	if (password === "")
		validationError = validationError + "Password is blank ! Please enter password.\n";
	if(validationError === "")
	{
		var url = host + "/login";
		var request = "action=LIN&username=" + username	+ "&password=" + password+ "&language=" + language;
		console.log("Sending request: " + request);
		$.post(url, request, function(response)
		{
			console.log(response);
			var resp = response.substr(0, response.indexOf(':'));

			console.log("response = " + resp);
			if(resp == "OK")
			{
				if(document.getElementById('rememberMe').checked)
				{
					console.log("Login user : " + document.getElementById('usernamelog').value + " saved");
					//var currentUser = document.getElementById('passwordlog').value;
					setCookie("return-user", document.getElementById('usernamelog').value, 365);
					// need to store password encrypted?
					//setCookie("return-login", currentLogin, 365);
				}
				else{
					setCookie("return-user", "", 365);
				}
				setCookie("current-user", document.getElementById('usernamelog').value, 1, "/");
				var status = response.substr(response.indexOf(':') + 1, 4);
				var redirect = response.substr(response.indexOf(':') + 6);
				//dialog_hide('divLogin');
				console.log("loaction.replace with: " + host + redirect);
				location.replace(host + redirect);
			}
			else if(resp == "ERROR")
			{
				var errorCode = response.substr(response.indexOf(':') + 1, 4);
				var errorString = response.substr(response.indexOf(':') + 6);
				alert("Failed to login...       Error Code : " + errorCode + "\n\n   " + errorString);
			}
			else
			{
				alert("Error: No response");
			}
		});
		//alert("Form Submitted Successfully...");
	}
	else
	{
		alert("Details are incomplete !\n\n" + validationError);
	}
}

function formLogout_submit()
{
	var language = currentLanguage;
	console.log("language : " + language);
	var url = host + "/logout";
	var request = "action=LOUT";
	console.log("Sending request: " + request);
	$.post(url, request, function(response)
	{
		console.log(response);
		var resp = response.substr(0, response.indexOf(':'));

		console.log("response = " + resp);
		if(resp == "OK")
		{
			var status = response.substr(response.indexOf(':') + 1, 4);
			var redirect = response.substr(response.indexOf(':') + 6);
			//dialog_hide('divLogin');
			console.log("loaction.replace with: " + host + redirect);
			location.replace(host + redirect);
		}
		else if(resp == "ERROR")
		{
			var errorCode = response.substr(response.indexOf(':') + 1, 4);
			var errorString = response.substr(response.indexOf(':') + 6);
			alert("Failed to login...       Error Code : " + errorCode + "\n\n   " + errorString);
		}
		else
		{
			alert("Error: No response");
		}
	});
	//alert("Form Submitted Successfully...");
}


function formNewPost_submit()
{
	if (document.getElementById('title').value == "" || document.getElementById('content').value == "")
	{
		alert("Fill All Fields !");
	}
	else
	{
		document.getElementById('formNewPost').submit();
		alert("Form Submitted Successfully...");
	}
}

function formNewClub_submit()
{
	if (document.getElementById('clubName').value == "" || document.getElementById('clubWebsite').value == "")
	{
		alert("Fill All Fields !");
	}
	else
	{
		document.getElementById('formNewClub').submit();
		alert("Form Submitted Successfully...");
	}
}
function createAccount_submit()
{
	dialog_hide('divLogin');
	dialog_show('divRegister');
}
function forgotPassword_submit()
{
	dialog_hide('divLogin');
	// dialog_show('divGetPassword')
}
//Function To Display Popup
function dialog_show(item, page)
{
	document.getElementById(item).style.display = "block";
}
//Function to Hide Popup
function dialog_hide(item)
{
	document.getElementById(item).style.display = "none";
}
//Function To Display Login Popup
function dialogLogin_show(item)
{
	var currentUser = "";
	currentUser = getCookie("return-user");
	console.log("Login user : " + currentUser);
	//createLoginDialog();
	if(currentUser != "")
	{
		document.getElementById('usernamelog').value = currentUser;
		document.getElementById('rememberMe').checked = true;
	}
	document.getElementById(item).style.display = "block";
}
