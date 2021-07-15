
# Simple-login-api

Simple login register api with refrance

# Routes

 - [ ] app.Post("/api/register", controllers.Register)
 - [ ] app.Post("/api/login", controllers.Login)
 - [ ] app.Get("/api/user", controllers.User)
 - [ ] app.Post("/api/logout", controllers.Logout)
 - [ ] app.Post("/api/change", controllers.ChangePassword)
 - [ ] app.Post("/api/send", controllers.SendPoint)
 - [ ] app.Get("/api/point", controllers.Point)
 - [ ] app.Get("/api/:market/:withdraw", controllers.Market)
 - [ ] app.Get("/api/upgrade", controllers.Upgrade)
 - [ ] app.Get("/api/most", controllers.MostUser)
 - [ ] app.Post("/api/check", controllers.EmailVerification)
 - [ ] app.Post("/api/setcheck", controllers.CheckEmailVerification)
 - [ ] app.Post("/api/forget", controllers.ForgetPassword)
 - [ ] app.Post("/api/setforget", controllers.ForgetChange)

# Doc

#**Register**

![enter image description here](https://i.hizliresim.com/q0k74v2.png)

Params :
{
"name":"username",
"email":"nikiyi5673@nhmty.com",
"password":"12345678",
"referance":"username" //choice
}

#**Email verification  code send**

![enter image description here](https://i.ibb.co/StzdZsB/Ekran-g-r-nt-s-2021-07-15-08-32-01.png)
![enter image description here](https://i.ibb.co/1sZz8Fs/Ekran-g-r-nt-s-2021-07-15-08-33-02.png)

Params : 
{
"email":"nikiyi5673@nhmty.com"
}




#**Check email verification  code**

![enter image description here](https://i.ibb.co/RY3XX8N/Ekran-g-r-nt-s-2021-07-15-08-34-01.png)

Params : 
{

"email":"nikiyi5673@nhmty.com",
"code":"YourCode"
}


#**Login**

![enter image description here](https://i.ibb.co/TBn3Jvr/Ekran-g-r-nt-s-2021-07-15-08-34-30.png)

Params :

{
"email":"nikiyi5673@nhmty.com",
"password":"12345678"
}

#**User**

//Return user info

![enter image description here](https://i.ibb.co/BjN6nVH/Ekran-g-r-nt-s-2021-07-15-08-35-04.png)

#**Point**

Get 2 point

![enter image description here](https://i.ibb.co/QXNfbSz/Ekran-g-r-nt-s-2021-07-15-08-35-27.png)

#**Market**

Buy item and send notification to email

![enter image description here](https://i.ibb.co/pjjRvBM/Ekran-g-r-nt-s-2021-07-15-08-52-07.png)

#**Send**

Send point to user

![enter image description here](https://i.ibb.co/64GK5qN/Ekran-g-r-nt-s-2021-07-15-08-50-57.png)

Params : 
{
"receiver":"receiver name",
"sendingPoints":"1"
}

#**Change Password**

![enter image description here](https://i.ibb.co/FwkVR1L/Ekran-g-r-nt-s-2021-07-15-08-58-23.png)

Params : 
{
"oldPassword":"12345678",
"newPassword":"abcde123321"
}

#**Logout**

![enter image description here](https://i.ibb.co/P9mSzWK/Ekran-g-r-nt-s-2021-07-15-09-00-22.png)

#**Check Upgrade**

Check upgrade from settings.ini

![enter image description here](https://i.ibb.co/vmkyC7V/Ekran-g-r-nt-s-2021-07-15-09-04-38.png)

#**Most Point Users**

List of most points users

![enter image description here](https://i.ibb.co/rvBdWdw/Ekran-g-r-nt-s-2021-07-15-09-05-40.png)

#**Send Forget Password Link**

![enter image description here](https://i.ibb.co/H2Bq7jW/Ekran-g-r-nt-s-2021-07-15-09-06-40.png)

Params : 
{
"email":"nikiyi5673@nhmty.com"
}

![enter image description here](https://i.ibb.co/nfy41WL/Ekran-g-r-nt-s-2021-07-15-09-07-43.png)

#**Check Forget Password Link**

![enter image description here](https://i.ibb.co/2g8MYtc/Ekran-g-r-nt-s-2021-07-15-09-09-14.png)

Params :
{

"token":"comming token from url params",
"newPassword":"new Password"

}
