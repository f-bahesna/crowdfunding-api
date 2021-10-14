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

### Day 3: October 02, 2021 

**Today's Progress**: Add Upload campaign image endpoint

**Thourghts:** Adding image with params [ is_primary ] utility to define, main image of campaign.

### Day 4: October 13, 2021

**Today's Progress**: Fix bugs upload campaign images endpoint

**Thourghts:** Struggled when uploading image filepath to database, because some parts in the database that Relation between campaign & campaignImages I've already set and that's the problem. I deleted relation and running as usual. not recommended but let's finish this course.

### Day 5: October 14, 2021 

**Today's Progress**: Fix input body params for is_primary campaign images

**Thourghts:** Deleted binding for is_primary input because that is boolean and when u want to set false just dont fill in the params.

