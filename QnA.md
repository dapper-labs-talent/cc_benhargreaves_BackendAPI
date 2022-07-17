How long did this assignment take?
> Around 8 or 9 hours across a few days

What was the hardest part?
> Getting Docker compose set up properly, and making sure the containers could talk to one another.

Did you learn anything new?
> Building JWT authentication with Golang! I've used auth mechanisms in the past with other languages (Oauth2 with both C# and NodeJS), but nothing in Go.

Is there anything you would have liked to implement but didn't have the time to?
> Unit Testing. With more time I would like to have started with the testing first but I got ahead of myself.
> Also, I would like to have implemented more robust error handling. A lot of the errors I'm returning to the user are fairly generic - I would like to have implemented more helpful error messaging to the end user

What are the security holes (if any) in your system? If there are any, how would you fix them?
> The first security hole is that I currently dont have a mechanism for invalidating the JWT. They do have an expiration of an hour, but in a real setting if the user was deleted/deactivated - There should be a way to make sure the JWT is no longer able to authenticate requests.

>There are a few ways to fix this - the quickest and dirtiest way would just be to change the secret key used to sign the tokens and then re-deploy. This would result in any existing tokens no longer being able to pass validation as their signature would now be incorrect. This comes with the pretty major drawback of invalidating ALL users tokens at once though.

> The preferable method would be to set up A server-side cache containing a blacklist of JWTs that have not yet expired, but have been revoked. A Request would then only be considered authenticated if the JWT was both valid, and not in the blacklist. This cache would need to be pruned regularly for JWTs that have since also expired

Do you feel that your skills were well tested?
> I think so - This certainly took some time going through package docs on my end for things I was unfamiliar with (like the JWT package)