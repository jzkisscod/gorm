# gorm
Gorm sample code
This is gorm sample code.
It will give sample for how to use gorm with belong to, one to many and multiple relation save and search

1. database connection<br/>
2. turn on debug<br/>
    There are two ways turning on GORM debug.<br/>
    2.1. Turn on debug on specific method:  db.Debug().Create<br/>
    2.2. Turn on log on all db operation: db.LogMode(true)<br/>
3. In one to many relation entities, As long as foreign key is defined, related entities can be saved at same time.<br/>
   But there is tricky in the saving associated entity.<br/>
   save associated entities<br/>
   3.1. when associated entities' id (primary key is set in advance), gorm will do the following operations on the
        associated entities<br/>
        it will do update, select and then insert<br/>
   3.2. when there is no id set, gorm will only do the insert.<br/>
   If performance is your concern, suggest not to set foreign key in gorm.<br/>
   If overwritten is your app's requirement, suggest to use foreign key in gorm.<br/>
4. Preload associate entities in one to many relation<br/>
5. Delete the entity without setting cascade relation.<br/>



