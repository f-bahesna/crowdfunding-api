# Days Of Progress - Log Recap

### Day 0: September 20, 2021
##### (I didn't count before)
#### There's a feature that I've made before:
- Register User endpoint
- Login endpoint
- Email check endpoint
- Avatar upload endpoint
- JWT Authentication
- Authentication Middleware
- List campaign endpoint
- Campaign detail endpoint


**Today's Progress**: Add create campaign endpoint

**Thourghts:** struggled with value in gin.Context Set func that I've been made,repeatedly I found an error for the mismatch of the value I took to retrieve the user id.for the rest nothing makes it difficult just want to understand better.

### Day 1: September 21, 2021

**Today's Progress**: Add update campaign endpoint

**Thourghts:** When get struct for input payload in handler, i got an error thats I wrote the wrong variable name. But its has solved. Ready to the next chapter.

### Day 2: September 23, 2021

**Today's Progress**: Add Upload campaign image endpoint

**Thourghts:** Adding image with params [ is_primary ] utility to define, main image of campaign.

### Day 3: October 13, 2021

**Today's Progress**: Fix bugs upload campaign images endpoint

**Thourghts:** Struggled when uploading image filepath to database, because some parts in the database that Relation between campaign & campaignImages I've already set and that's the problem. I deleted relation and running as usual. not recommended but let's finish this course.

### Day 4: October 14, 2021 

**Today's Progress**: Fix input body params for is_primary campaign images & Prepare for transaction service

**Thourght: 1.**  Deleted binding for is_primary input because that is boolean and when u want to set false just dont fill in the params.

**Thourght: 2.** Create domain dan entity for transation service.  

### Day 5: October 15, 2021 

**Today's Progress**:
- Create new endpoint campaign transaction list by campaign id.
- Authorization for campaign transaction list by user who created campaign.

**Thourght: 1.**  Deleted binding for is_primary input because that is boolean and when u want to set false just dont fill in the params.

**Concern** Authorization succeeds gives an error if the campaign transaction list isn't the user who made it. but if more than one and there is another user_id in transaction list , the transaction still appears.

### Day 6: October 21, 2021 

**Progress**: Add Transaction Endpoint

**Thourght:** Add Transaction handler , Service, Repository, Entity and the last build formatter json response for API response. 

### Day 7: November 29, 2021 

**Progress**: Add Docker Configuration With Postgresql Database

**Thourght:** I've so many make mistake when setup docker in here, such as wrong configuration in dockerfile, wrong host for connection go container with postgre database because can't connect localhost from Docker and I solve with "docker.for.mac.localhost" in main.go file and finally works!



