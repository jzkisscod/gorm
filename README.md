# gorm
Gorm sample code
This is gorm sample code.
It will give sample for how to use gorm with belong to, one to many and multiple relation save and search

1. database connection
2. turn on debug
    There are two ways turning on GORM debug.
    1.a. Turn on debug on specific method:  db.Debug().Create
    1.b. Turn on log on all db operation: db.LogMode(true)
3. In one to many relation entities, As long as foreign key is defined, related entities can be saved at same time.
   But there is tricky in the saving associated entity.
   save associated entities
   3.a. when associated entities' id (primary key is set in advance), gorm will do the following operations on the
        associated entities
        it will do update, select and then insert
   3.b. when there is no id set, gorm will only do the insert.
   If performance is your concern, suggest not to set foreign key in gorm.
   If overwritten is your app's requirement, suggest to use foreign key in gorm.
4. Preload associate entities in one to many relation
5. Delete the entity without setting cascade relation.



