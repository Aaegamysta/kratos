# My Solution and Attempt
## An Ideal Solution
### User Behavior Analytics
- Having a sort of User Behavior Analytics as 3rd party service provide would definetly help us analyzing the collected sessions such as unusual login times.
### Provide a way for TOR detection
- This is might be out of scope, but it is also important to keep in mind that there are some people who using TOR to attempt logins with harmful intent.
- Various third-party APIs, such as ip2tor, allow you to check whether an IP is associated with a TOR exit node.
- TOR Project’s Exit Node List: The TOR Project maintains a list of exit nodes. You can periodically download this list and check IPs against it.
### Having a Machine Learning Model
- It is common these days to stick AI with everything, but in this instance, a machine learning model would be really helpful to help decide what kind of login might be noise or actually dangerous.
- This machine learning model can be trained on the various session data that will be saved and accumulated over time.
## Solution
### Shortcircuiting the Algorithm to Make It Faster
- There are multiple factors that were taken into account to try make the algorithm faster and prevent checking all past sessions and this is known as shortcircuiting.
- Leveraging these factors can be found in the implementation and it is fully documented in manager_http.go
### Why Http Manager?
- I have chosen to add a private method in Http Manager because all it is the point where session management is done as it implements SessionManager interface.
- I have decided for this function to be part a private of the implementation of Http Manager and not part of the SessionManager interface because it does not make much sense to export it to the outer world and outside the package.
- This method exclusively gets run when attempting to *activate the session* which in turns gets called by different login flows.
- I have not decided to make it part of the Login flow, as it would be more useful to make it more abstract and leave it in the Session and HttpManager component because it might be
used later on not just in login flow.
- Applying the solution directly to login flow would be brittle.
## Shorting Comings, Flaws and Possible Improvements
### GeoLocation Being Part of Header
- The location attribute in Devices was just the city, region and the country (coming from http headers). A better approach would have to acquire the geolocation in terms of latitude and longitude.
- This can be done via reliable services such as www.ipinfo.io
- This will still have its short comings because there are some services that will not always provide accurate and perfect result.
- A better approach would to keep the location as it is and two extra attributes for geolocation: latitude and longitude that the algorithm will mainly depend on.
### Testing Last Path in flagSession
- Because of time constraints, I could not test the function where it needs to check the past **N** inactive sessions and see if there is a problem with the login attempt
and if the are are far given the time window.
### Relying on old methods to create sessions 
- There should have been a better way to manage the lifecycle of sessions in test and not just be only dependent on requests for them to be created.
- There should have been a better way to mock them and provide different factory functions to be fitting the for the test case.
- flagSession should have been tested separately, but the problem was that ActivateSession method is an integral part setting up the Session for testing.
### Integration and Benchmark tests
- There should have been more integration tests that rely on **ActivateSession** method such as 
- There also should have been tests that bench mark the configuration and see if they meet the timing constraints.
- Load tests are also important to configure the parameters that the algorithms depend on.
## Possible Improvements
### IP Address Ranges
- A good approach to provide a shortcircuit is to determing from IP if it belongs to a cached IP ranges of the user or ip ranges of the country/region that he is in.
- Another approach is also for determining the IP ranges and CIDR Block of the user and to calculate it. This can be done without relying on external services.
- Another possible improvement would also be relying on a service that tells us if the incoming IP is from a VPN or a proxy
### Unusual Geographic Locations and Locations that are source repeated cyber attacks
- There should be a check for IP ranges of countries that history of conducting cyber attacks or cyber ransomware such as Russia, China or North Korea.
### User Agent
- Monitoring the user agent and using libraries such as FingerprintJS that uniquely assign a fingerprint to the device of the user can help spoof the.
- Monitoring the user agent however requires some testdata to test an algorithm that would detect a pattern in the change of the user agents, which is not availabe and I think lies outsiide the scope of the task.
- The data that can be provided by Fingerprint JS must be delivered via the browser and saved accordingly. Because of the lack of a way to do this (since frontend is not available), it could not be done but might be a good suggestion as an addition.
### Prevention of Noise
- The configuration should not be strict, but rather also introduce levels of that configuration that would mark the risk low, medium or high.
- For the speed, for example, there should not only be one speed that we will be comparing to and deciding if the risk is **high** or **none** (binary), but rather there should be 
an intermediate speed that would should mark it as **suspicious**.
- This middle level marking will spare us from bombarding the user with notifications and not to ignore it completely and could be helpful when it comes to aggregating and analyzing the data later on.
### Making the Algorithm as much as Customizable as Possible
- The flow of flagging the sessions could be modelled in a pipeline where each stage represents a check.
- It is possible to make the algorithm stricter by adding more pipeline stages or removing them if there is too much noise (depending on the client's wish)
- This emulates the behavior of a plugin that can be applied to customize the behavior of the software.

## Final Words
I would like to thank you and I am looking forward for the exciting discussion that I could be having with you and the team regarding these points and thoughts.