OktaCSV pulls the Okta Event logs and streams the output to the Console.

It is written in the GO Language, and can run interpreted (GO Compiles it on the fly),
or run the contained binaries for the platform you are running.

It is very easy to Modify and Customize, and I hope others find it useful


c:\oktacsv.exe https://yourORG.okta.com 0YOURKEYHERE-9UW4H9r6j6scgnzbzx0bePfRz 
OktaCSV by Patrick McDowell pmcdowell@okta.com

   ___  _   _         ___ _____   __
  / _ \| |_| |_ __ _ / __/ __\ \ / /
 | (_) | / /  _/ _` | (__\__ \\ V /
  \___/|_\_\\__\__,_|\___|___/ \_/

OktaCSV is checking to see what time it is in OktaLand, and starting to follow the System Log
This can take a few seconds, but it is worth the wait
2016-12-05T21:26:02.000Z , Successfully issued [access_token - id: AT.w3DoooWIbqdK_DxXCuisGxxcVBNYAW1E1-V7
2016-12-05T21:26:02.000Z , Successfully issued [refresh_token - id: oar1zn0j832uHfEN20h6, hash: bFmx7iRQAQ
2016-12-05T21:26:03.000Z , Successfully issued [access_token - id: AT.OnMxvvfht03FIx1Onw7Uj_HDJwtui5_fQ6Ig
2016-12-05T21:26:03.000Z , Successfully issued [refresh_token - id: oar1zmozmgisfcK6d0h6, hash: PeLXVNjiJ2
2016-12-05T21:26:04.000Z , Successfully issued [access_token - id: AT.qmtNQzHbVif9e8WWHyqoDuQfU60DpULvrnth
2016-12-05T21:26:04.000Z , Successfully issued [refresh_token - id: oar1zmr9wQ4Sq1K5k0h6, hash: zocwyxck8y

...


OktaCSV also runs Great in Docker !!

C:\temp\go\oktacsv>docker run -v c:\temp\go\oktacsv:/oktacsv golang bash -c "go run /oktacsv/oktacsv.go https://customers.oktapreview.com 00eRP0hnRFqQE8JGw9UW4H9r6j6scgnzbzx0bePfRz"

