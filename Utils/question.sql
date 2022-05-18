Create Database GodAge;
use GodAge;
create user GodAger identified by 'qs123456';
grant select, update On GodAge.* to GodAger;
Create Table QuestionList
(
    QuestionId   int auto_increment primary key,
    QuestionType smallint not null default 0,
    CategoryId   int      not null default 0
) charset = utf8mb4;
Create Index CategoryId On QuestionList (CategoryId);

Create Table QuestionInfo
(
    QuestionId     int primary key,
    CreateDate     int          default 0,
    LastChangeDate int          default 0,
    Title          varchar(255) default '空',
    TitleHasImage  smallint     default 0
) charset = utf8mb4;

Create Table Options
(
    OptionId    int auto_increment primary key,
    QuestionId  int,
    OptionTitle varchar(255)
) charset = utf8mb4;

Create Table User
(
    UserId     int auto_increment primary key,
    StudentNum int unique,
    Password   varchar(255)
) charset = utf8mb4;

Create Table UserInfo
(
    UserId     int primary key,
    RealName   varchar(10) default '狗东西没名字',
    NickName   varchar(10) default '舔狗',
    AttendDate int         default 0,
    Major      varchar(10) default '未选择',
    Birthday   int         default 0,
    Phone      int         default -1,
    Email      varchar(30) default '空'
) charset = utf8mb4;

Create Table QuestionsFinish
(
    UserId          int,
    QuestionId      int,
    primary key (UserId, QuestionId),
    FirstFinishDate int      default 0,
    LastFinishDate  int      default 0,
    isCorrect       smallint default -1
) charset = utf8mb4;
Create Index IdDate On QuestionsFinish (UserId, FirstFinishDate) using Btree;

Create Table Subject
(
    SubjectId   int auto_increment primary key,
    SubjectName varchar(30) default '空'
) charset = utf8mb4,
  engine = MyISAM;
Create FULLTEXT INDEX SearchSubject On Subject (SubjectName);

Create Table Section
(
    SectionId   int auto_increment primary key,
    SubjectId   int         default -1,
    SectionName varchar(40) default '空'
) charset = utf8mb4, engine = MyISAM;
Create FULLTEXT INDEX SearchSection On Section (SectionName);
Create Index SearchBySubject On Section (SubjectId);

Create Table Lesson
(
    LessonId   int auto_increment primary key,
    SectionId  int         default -1,
    LessonName varchar(40) default '空'
) charset = utf8mb4, engine = MyISAM;
Create FULLTEXT INDEX SearchLesson On Lesson (LessonName);
Create Index SearchBySection On Lesson (SectionId);

Create Table UserLesson
(
    UserId    int,
    SubjectId int,
    primary key (UserId, SubjectId)
);