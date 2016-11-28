function tr(text, code)
{
	//console.log("tr" + "." + code + "." + text);
	var translatedText = translator[code][text];
	//console.log("translatedText = (" + translatedText + ")");
	if(translatedText == undefined)
		return text;
	return translatedText;
}

var translator =
{
	"ge": {
		"and join a club or group to get the most from the page.": "und einem Verein beitreten oder eine Gruppe, die meisten von der Seite zu gelangen.",
		"Add New Events": "Fugen Neue Geschehen",
		"Add New Results": "Fugen Neue Ergebnisse",
		"Administrator": "Verwalter",
		"Advertisement": "Anzeige",
		"Ammend Events": "Andern Événements",
		"Ammend Post": "Andern Nachricht",
		"Ammend Results": "Andern Geschehen",
		"Club Page": "Clubseite",
		"Create An Account": "Ein Konto Erstellen",
		"Create New Post": "Erstellen Neue Nachricht",
		"Events": "Geschehen",
		"Forgot Your Password?": "Haben Sie Ihr Passwort Vergessen?",
		"Hello": "Hallo",
		"Home Page": "Startseite",
		"Join Page": "Beizutreten Seite",
		"Login": "Login",
		"LOGIN..... ": "LOGIN..... ",
		"Manage Club Details": "Verwalter Verein Informationen",
		"Manage Club Sessions": "Verwalter Verein Sitzungen",
		"Manage Events Page": "Verwalter Événements Seite",
		"Manage Group Sessions": "Verwalter Gruppe Sitzungen",
		"Manage Results Page": "Verwalter Ergebnisse Seite",
		"Moderator": "Mäßigen",
		"Password": "Passwort",
		"Please Login to your account": "Bitte Login auf Ihr Konto",
		"Remember Me": "Erinnere Dich an Mich",
		"Results": "Ergebnisse",
		"+ Register ": "+ Registrieren  ",
		"Search Preferences": "Sucheinstellungen",
		"Username": "Benutzername"
		},
	"fr": {
		"and join a club or group to get the most from the page.": "et adhérer à un club ou d'un groupe pour tirer le meilleur de la page.",
		"Add New Events": "Adjouter Nouveaux Événements",
		"Add New Results": "Adjouter Nouveaux Résultats",
		"Administrator": "Administrateur",
		"Advertisement": "Publicité",
		"Ammend Events": "Modifier Événements",
		"Ammend Post": "Modifier Message",
		"Ammend Results": "Modifier Résultats",
		"Club Page": "Club Page",
		"Create An Account": "Créer un Compte",
		"Create New Post": "Créer Nouveau Message",
		"Events": "Événements",
		"Forgot Your Password?": "Mot de Passe Oublié?",
		"Hello": "Bonjour",
		"Home Page": "Accueil",
		"Join Page": "Joignez Page",
		"Login": "Connexion",
		"LOGIN..... ": "CONNEXION..... ",
		"Manage Club Details": "Faire Club Détails",
		"Manage Club Sessions": "Faire Club Sessions",
		"Manage Events Page": "Faire Événements Page",
		"Manage Group Sessions": "Faire Groupe Sessions",
		"Manage Results Page": "Faire Résultats Page",
		"Moderator": "Modérateur",
		"Password": "Mot de Passe",
		"Please Login to your account": "S'il vous plaît vous connecter à votre compte",
		"Remember Me": "Souviens-Toi de Moi",
		"Results": "Résultats",
		"+ Register ": "+ Registre  ",
		"Search Preferences": "Recherche Préférences",
		"Username": "Nom d'utilisateur"
		},
	"en": {
		"and join a club or group to get the most from the page.": "and join a club or group to get the most from the page.",
		"Add New Events": "Add New Events",
		"Add New Results": "Add New Results",
		"Administrator": "Administrator",
		"Advertisement": "Advertisement",
		"Ammend Events": "Ammend Events",
		"Ammend Post": "Ammend Post",
		"Ammend Results": "Ammend Results",
		"Club Page": "Club Page",
		"Create An Account": "Create An Account",
		"Create New Post": "Create New Post",
		"Events": "Events",
		"Forgot Your Password?": "Forgot Your Password?",
		"Hello": "Hello",
		"Home Page": "Home Page",
		"Join Page": "Join Page",
		"Login": "Login",
		"LOGIN..... ": "LOGIN..... ",
		"Manage Club Details": "Manage Club Details",
		"Manage Club Sessions": "Manage Club Sessions",
		"Manage Events Page": "Manage Events Page",
		"Manage Group Sessions": "Manage Group Sessions",
		"Manage Results Page": "Manage Results Page",
		"Moderator": "Moderator",
		"Password": "Password",
		"Please Login to your account": "Please Login to your account",
		"Remember Me": "Remember Me",
		"Results": "Results",
		"+ Register ": "+ Register ",
		"Search Preferences": "Search Preferences",
		"Username": "Username"
		}
};


