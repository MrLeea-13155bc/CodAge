Create Database GodAge;
use GodAge;
create user GodAger identified by 'qs123456';
grant all On GodAge.* to GodAger;
Create Table QuestionList
(
    QuestionId   int auto_increment primary key,
    QuestionType smallint not null default 0,
    SectionId   int      not null default 0
) charset = utf8mb4;
Create Index CategoryId On QuestionList (SectionId);

Create Table QuestionInfo
(
    QuestionId     int primary key,
    CreateDate     int          default 0,
    LastChangeDate int          default 0,
    Title          varchar(255) default '空',
    CorrectAnswer  varchar(30) default '',
    ImageUrl  varchar(255) default ''
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
    RealName   varchar(10)  default '狗东西没名字',
    AttendDate int          default 0,
    Major      varchar(10)  default '未选择',
    Icon       varchar(255) default 'nil'
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

Create Table UserLesson
(
    UserId    int,
    SubjectId int,
    primary key (UserId, SubjectId)
);


insert Into Subject Set SubjectName = '舔狗学';
INSERT INTO Section VALUES (1,1, '大气概况');
INSERT INTO Section VALUES (2,1, '气温');
INSERT INTO Section VALUES (3,1, '气压');
INSERT INTO Section VALUES (4,1, '大气水平运动');
INSERT INTO Section VALUES (5,1, '大气环流');
INSERT INTO Section VALUES (6,1, '大气湿度和稳定度');
INSERT INTO Section VALUES (7,1, '云和降水');
INSERT INTO Section VALUES (8,1, '雾和能见度');
INSERT INTO Section VALUES (9,1, '海流');
INSERT INTO Section VALUES (10,1, '海浪');
INSERT INTO Section VALUES (11,1, '气团和锋');
INSERT INTO Section VALUES (12,1, '锋面气旋');
INSERT INTO Section VALUES (13,1, '冷高压');
INSERT INTO Section VALUES (14,1, '副热带高压');
INSERT INTO Section VALUES (15,1, '热带气旋');
INSERT INTO QuestionList VALUES (1, 0, 1);
INSERT INTO QuestionList VALUES (2, 0, 1);
INSERT INTO QuestionList VALUES (3, 0, 1);
INSERT INTO QuestionList VALUES (4, 0, 1);
INSERT INTO QuestionList VALUES (5, 0, 1);
INSERT INTO QuestionList VALUES (6, 0, 1);
INSERT INTO QuestionList VALUES (7, 0, 1);
INSERT INTO QuestionList VALUES (8, 0, 1);
INSERT INTO QuestionList VALUES (9, 0, 1);
INSERT INTO QuestionList VALUES (10, 0, 1);
INSERT INTO QuestionList VALUES (11, 0, 1);
INSERT INTO QuestionList VALUES (12, 0, 1);
INSERT INTO QuestionList VALUES (13, 0, 1);
INSERT INTO QuestionList VALUES (14, 0, 1);
INSERT INTO QuestionList VALUES (15, 0, 1);
INSERT INTO QuestionList VALUES (16, 0, 1);
INSERT INTO QuestionList VALUES (17, 0, 1);
INSERT INTO QuestionList VALUES (18, 0, 1);
INSERT INTO QuestionList VALUES (19, 0, 1);
INSERT INTO QuestionList VALUES (20, 0, 1);
INSERT INTO QuestionList VALUES (21, 0, 1);
INSERT INTO QuestionList VALUES (22, 0, 1);
INSERT INTO QuestionList VALUES (23, 0, 1);
INSERT INTO QuestionList VALUES (24, 0, 1);
INSERT INTO QuestionList VALUES (25, 0, 1);
INSERT INTO QuestionList VALUES (26, 0, 1);
INSERT INTO QuestionList VALUES (27, 1, 1);
INSERT INTO QuestionList VALUES (28, 1, 1);
INSERT INTO QuestionList VALUES (29, 1, 1);
INSERT INTO QuestionList VALUES (30, 1, 1);
INSERT INTO QuestionList VALUES (31, 1, 1);


Insert Into QuestionInfo Set CorrectAnswer= '4',QuestionId= 1 ,Title=' 大气中具有透过太阳短波辐射、吸收地面长波辐射的主要气体 ';
Insert Into QuestionInfo Set CorrectAnswer= '5',QuestionId= 2 ,Title=' 在大气成分中,主要吸收太阳紫外线的气体成分为: ';
Insert Into QuestionInfo Set CorrectAnswer= '11',QuestionId= 3 ,Title=' 在自然界的温度和压力条件下，哪种大气成份能在气态、液态和固态三者之间互相转化? ';
Insert Into QuestionInfo Set CorrectAnswer= '16',QuestionId= 4 ,Title=' 在下列大气成分中,能产生温室效应的大气成分是: ';
Insert Into QuestionInfo Set CorrectAnswer= '19',QuestionId= 5 ,Title=' 大气密度随着高度的增加而______ ';
Insert Into QuestionInfo Set CorrectAnswer= '22',QuestionId= 6 ,Title=' 在气压相同的情况下,哪种特征的空气其密度最小? ';
Insert Into QuestionInfo Set CorrectAnswer= '27',QuestionId= 7 ,Title=' 在气压相同的情况下,密度较大的空气是: ';
Insert Into QuestionInfo Set CorrectAnswer= '30',QuestionId= 8 ,Title=' 地球大气最低层称为对流层，其平均厚度约为: ';
Insert Into QuestionInfo Set CorrectAnswer= '33',QuestionId= 9 ,Title=' 对流层的厚度随纬度有较大的变化，其最厚出现在: ';
Insert Into QuestionInfo Set CorrectAnswer= '38',QuestionId= 10 ,Title=' 对流层的厚度随季节变化，哪个季节最厚? ';
Insert Into QuestionInfo Set CorrectAnswer= '44',QuestionId= 11 ,Title=' 对流层的厚度随季节变化，其最薄出现在: ';
Insert Into QuestionInfo Set CorrectAnswer= '48',QuestionId= 12 ,Title=' 对流层具有的特点之一是: ';
Insert Into QuestionInfo Set CorrectAnswer= '50',QuestionId= 13 ,Title=' 一般来说，在对流层中平均每升高100米,气温下降约为: ';
Insert Into QuestionInfo Set CorrectAnswer= '56',QuestionId= 14 ,Title=' 在对流层中,气温垂直递减率约为: ';
Insert Into QuestionInfo Set CorrectAnswer= '57',QuestionId= 15 ,Title=' 在对流层内，温度和湿度: ';
Insert Into QuestionInfo Set CorrectAnswer= '63',QuestionId= 16 ,Title=' 大气中主要的天气现象发生在: ';
Insert Into QuestionInfo Set CorrectAnswer= '65',QuestionId= 17 ,Title=' 云、雾、雨、雪等大气中的主要天气现象都发生在: ';
Insert Into QuestionInfo Set CorrectAnswer= '70',QuestionId= 18 ,Title=' 在对流层中气温通常随着高度的降低而 ';
Insert Into QuestionInfo Set CorrectAnswer= '73',QuestionId= 19 ,Title=' 在对流层中气温涌常随着高度的升高而 ';
Insert Into QuestionInfo Set CorrectAnswer= '80',QuestionId= 20 ,Title=' 自由大气的起始高度大约为 ';
Insert Into QuestionInfo Set CorrectAnswer= '84',QuestionId= 21 ,Title=' 根据对流层中______的不同特征，可将其分为摩擦层和自由大气两个层次。 ';
Insert Into QuestionInfo Set CorrectAnswer= '86',QuestionId= 22 ,Title=' 对流层可分为摩擦层和自由大气两层,摩擦层的平均厚度为: ';
Insert Into QuestionInfo Set CorrectAnswer= '91',QuestionId= 23 ,Title=' 表征大气状态的物理量和物理现象的气象术语称为: ';
Insert Into QuestionInfo Set CorrectAnswer= '93',QuestionId= 24 ,Title=' 下列哪组完全属于气象要素 ';
Insert Into QuestionInfo Set CorrectAnswer= '99',QuestionId= 25 ,Title=' 下列哪个等压面最能代表对流层大气的一般运动状况? ';
Insert Into QuestionInfo Set CorrectAnswer= '102',QuestionId= 26 ,Title=' 大气的垂直分层自上而下依次为 ';
Insert Into QuestionInfo Set CorrectAnswer= '',QuestionId= 27 ,Title=' 对大气温度分布有较大影响的大气成份: ';
Insert Into QuestionInfo Set CorrectAnswer= '',QuestionId= 28 ,Title=' 下列哪些成份能强烈地吸收和放射长波辐射，对地面和大气的温度有较大的影响? ';
Insert Into QuestionInfo Set CorrectAnswer= '',QuestionId= 29 ,Title=' 对流层厚度随着下列哪些因素变化? ';
Insert Into QuestionInfo Set CorrectAnswer= '',QuestionId= 30 ,Title=' 对流层的主要特征 ';
Insert Into QuestionInfo Set CorrectAnswer= '',QuestionId= 31 ,Title=' 在气压不变的情况下，下列哪些条件可使空气密度变小? ';

Insert Into Options Set OptionId = 1 ,OptionTitle=  '臭氧' ,QuestionId=  1 ;
Insert Into Options Set OptionId = 2 ,OptionTitle=  '氮气' ,QuestionId=  1 ;
Insert Into Options Set OptionId = 3 ,OptionTitle=  '氧气' ,QuestionId=  1 ;
Insert Into Options Set OptionId = 4 ,OptionTitle=  '二氧化碳' ,QuestionId=  1 ;
Insert Into Options Set OptionId = 5 ,OptionTitle=  '臭氧' ,QuestionId=  2 ;
Insert Into Options Set OptionId = 6 ,OptionTitle=  '二氧化碳' ,QuestionId=  2 ;
Insert Into Options Set OptionId = 7 ,OptionTitle=  '氧气' ,QuestionId=  2 ;
Insert Into Options Set OptionId = 8 ,OptionTitle=  '氮气' ,QuestionId=  2 ;
Insert Into Options Set OptionId = 9 ,OptionTitle=  '氢气' ,QuestionId=  3 ;
Insert Into Options Set OptionId = 10 ,OptionTitle=  '氧气' ,QuestionId=  3 ;
Insert Into Options Set OptionId = 11 ,OptionTitle=  '水汽' ,QuestionId=  3 ;
Insert Into Options Set OptionId = 12 ,OptionTitle=  '二氧化碳' ,QuestionId=  3 ;
Insert Into Options Set OptionId = 13 ,OptionTitle=  '氧气' ,QuestionId=  4 ;
Insert Into Options Set OptionId = 14 ,OptionTitle=  '水汽' ,QuestionId=  4 ;
Insert Into Options Set OptionId = 15 ,OptionTitle=  '臭氧' ,QuestionId=  4 ;
Insert Into Options Set OptionId = 16 ,OptionTitle=  '二氧化碳' ,QuestionId=  4 ;
Insert Into Options Set OptionId = 17 ,OptionTitle=  '迅速增加' ,QuestionId=  5 ;
Insert Into Options Set OptionId = 18 ,OptionTitle=  '缓慢增加' ,QuestionId=  5 ;
Insert Into Options Set OptionId = 19 ,OptionTitle=  '迅速减小' ,QuestionId=  5 ;
Insert Into Options Set OptionId = 20 ,OptionTitle=  '不变' ,QuestionId=  5 ;
Insert Into Options Set OptionId = 21 ,OptionTitle=  '暖干' ,QuestionId=  6 ;
Insert Into Options Set OptionId = 22 ,OptionTitle=  '暖湿' ,QuestionId=  6 ;
Insert Into Options Set OptionId = 23 ,OptionTitle=  '冷干' ,QuestionId=  6 ;
Insert Into Options Set OptionId = 24 ,OptionTitle=  '冷湿' ,QuestionId=  6 ;
Insert Into Options Set OptionId = 25 ,OptionTitle=  '暖干' ,QuestionId=  7 ;
Insert Into Options Set OptionId = 26 ,OptionTitle=  '暖湿' ,QuestionId=  7 ;
Insert Into Options Set OptionId = 27 ,OptionTitle=  '冷干' ,QuestionId=  7 ;
Insert Into Options Set OptionId = 28 ,OptionTitle=  '冷湿' ,QuestionId=  7 ;
Insert Into Options Set OptionId = 29 ,OptionTitle=  '1～2km' ,QuestionId=  8 ;
Insert Into Options Set OptionId = 30 ,OptionTitle=  '10～12km' ,QuestionId=  8 ;
Insert Into Options Set OptionId = 31 ,OptionTitle=  '6～8km' ,QuestionId=  8 ;
Insert Into Options Set OptionId = 32 ,OptionTitle=  '17～18km' ,QuestionId=  8 ;
Insert Into Options Set OptionId = 33 ,OptionTitle=  '赤道低纬地区' ,QuestionId=  9 ;
Insert Into Options Set OptionId = 34 ,OptionTitle=  '中纬度地区' ,QuestionId=  9 ;
Insert Into Options Set OptionId = 35 ,OptionTitle=  '高纬度地区' ,QuestionId=  9 ;
Insert Into Options Set OptionId = 36 ,OptionTitle=  '极地地区' ,QuestionId=  9 ;
Insert Into Options Set OptionId = 37 ,OptionTitle=  '春季' ,QuestionId=  10 ;
Insert Into Options Set OptionId = 38 ,OptionTitle=  '夏季' ,QuestionId=  10 ;
Insert Into Options Set OptionId = 39 ,OptionTitle=  '秋季' ,QuestionId=  10 ;
Insert Into Options Set OptionId = 40 ,OptionTitle=  '冬季' ,QuestionId=  10 ;
Insert Into Options Set OptionId = 41 ,OptionTitle=  '春季' ,QuestionId=  11 ;
Insert Into Options Set OptionId = 42 ,OptionTitle=  '夏季' ,QuestionId=  11 ;
Insert Into Options Set OptionId = 43 ,OptionTitle=  '秋季' ,QuestionId=  11 ;
Insert Into Options Set OptionId = 44 ,OptionTitle=  '冬季' ,QuestionId=  11 ;
Insert Into Options Set OptionId = 45 ,OptionTitle=  '空气不易产生平流运动' ,QuestionId=  12 ;
Insert Into Options Set OptionId = 46 ,OptionTitle=  '气象要素水平分布均匀' ,QuestionId=  12 ;
Insert Into Options Set OptionId = 47 ,OptionTitle=  '气温随高度增加' ,QuestionId=  12 ;
Insert Into Options Set OptionId = 48 ,OptionTitle=  '气象要素水平分布不均匀' ,QuestionId=  12 ;
Insert Into Options Set OptionId = 49 ,OptionTitle=  '0.8℃' ,QuestionId=  13 ;
Insert Into Options Set OptionId = 50 ,OptionTitle=  '0.65℃' ,QuestionId=  13 ;
Insert Into Options Set OptionId = 51 ,OptionTitle=  '0.3℃' ,QuestionId=  13 ;
Insert Into Options Set OptionId = 52 ,OptionTitle=  '0.25℃' ,QuestionId=  13 ;
Insert Into Options Set OptionId = 53 ,OptionTitle=  '1℃/100米' ,QuestionId=  14 ;
Insert Into Options Set OptionId = 54 ,OptionTitle=  '25℃/100米' ,QuestionId=  14 ;
Insert Into Options Set OptionId = 55 ,OptionTitle=  '5℃/100米' ,QuestionId=  14 ;
Insert Into Options Set OptionId = 56 ,OptionTitle=  '65℃/100米' ,QuestionId=  14 ;
Insert Into Options Set OptionId = 57 ,OptionTitle=  '沿水平方向分布不均匀' ,QuestionId=  15 ;
Insert Into Options Set OptionId = 58 ,OptionTitle=  '沿水平方向分布均匀' ,QuestionId=  15 ;
Insert Into Options Set OptionId = 59 ,OptionTitle=  '随高度的增加而增加' ,QuestionId=  15 ;
Insert Into Options Set OptionId = 60 ,OptionTitle=  '沿垂直方向分布均匀' ,QuestionId=  15 ;
Insert Into Options Set OptionId = 61 ,OptionTitle=  '对流层' ,QuestionId=  16 ;
Insert Into Options Set OptionId = 62 ,OptionTitle=  '平流层' ,QuestionId=  16 ;
Insert Into Options Set OptionId = 63 ,OptionTitle=  '热层' ,QuestionId=  16 ;
Insert Into Options Set OptionId = 64 ,OptionTitle=  '中间层' ,QuestionId=  16 ;
Insert Into Options Set OptionId = 65 ,OptionTitle=  '热层' ,QuestionId=  17 ;
Insert Into Options Set OptionId = 66 ,OptionTitle=  '平流层' ,QuestionId=  17 ;
Insert Into Options Set OptionId = 67 ,OptionTitle=  '对流层' ,QuestionId=  17 ;
Insert Into Options Set OptionId = 68 ,OptionTitle=  '中间层' ,QuestionId=  17 ;
Insert Into Options Set OptionId = 69 ,OptionTitle=  '下降' ,QuestionId=  18 ;
Insert Into Options Set OptionId = 70 ,OptionTitle=  '上升' ,QuestionId=  18 ;
Insert Into Options Set OptionId = 71 ,OptionTitle=  '先升后降' ,QuestionId=  18 ;
Insert Into Options Set OptionId = 72 ,OptionTitle=  '先降后升' ,QuestionId=  18 ;
Insert Into Options Set OptionId = 73 ,OptionTitle=  '下降' ,QuestionId=  19 ;
Insert Into Options Set OptionId = 74 ,OptionTitle=  '上升' ,QuestionId=  19 ;
Insert Into Options Set OptionId = 75 ,OptionTitle=  '先升后降' ,QuestionId=  19 ;
Insert Into Options Set OptionId = 76 ,OptionTitle=  '先降后升' ,QuestionId=  19 ;
Insert Into Options Set OptionId = 77 ,OptionTitle=  '7~8km' ,QuestionId=  20 ;
Insert Into Options Set OptionId = 78 ,OptionTitle=  '5～6km' ,QuestionId=  20 ;
Insert Into Options Set OptionId = 79 ,OptionTitle=  '3～4km' ,QuestionId=  20 ;
Insert Into Options Set OptionId = 80 ,OptionTitle=  '1～1.5km' ,QuestionId=  20 ;
Insert Into Options Set OptionId = 81 ,OptionTitle=  '气温' ,QuestionId=  21 ;
Insert Into Options Set OptionId = 82 ,OptionTitle=  '气压' ,QuestionId=  21 ;
Insert Into Options Set OptionId = 83 ,OptionTitle=  '湿度' ,QuestionId=  21 ;
Insert Into Options Set OptionId = 84 ,OptionTitle=  '大气运动' ,QuestionId=  21 ;
Insert Into Options Set OptionId = 85 ,OptionTitle=  '7～8km' ,QuestionId=  22 ;
Insert Into Options Set OptionId = 86 ,OptionTitle=  '5～6k' ,QuestionId=  22 ;
Insert Into Options Set OptionId = 87 ,OptionTitle=  '3～4km' ,QuestionId=  22 ;
Insert Into Options Set OptionId = 88 ,OptionTitle=  '1～1.5km' ,QuestionId=  22 ;
Insert Into Options Set OptionId = 89 ,OptionTitle=  '天气' ,QuestionId=  23 ;
Insert Into Options Set OptionId = 90 ,OptionTitle=  '气候' ,QuestionId=  23 ;
Insert Into Options Set OptionId = 91 ,OptionTitle=  '气象要素' ,QuestionId=  23 ;
Insert Into Options Set OptionId = 92 ,OptionTitle=  '天气系统' ,QuestionId=  23 ;
Insert Into Options Set OptionId = 93 ,OptionTitle=  '风、云、雾、霜、沙尘暴' ,QuestionId=  24 ;
Insert Into Options Set OptionId = 94 ,OptionTitle=  '气压、高气压、台风' ,QuestionId=  24 ;
Insert Into Options Set OptionId = 95 ,OptionTitle=  '风、云、雨、冷锋、暖锋' ,QuestionId=  24 ;
Insert Into Options Set OptionId = 96 ,OptionTitle=  '气温、气压、冷锋、暖锋' ,QuestionId=  24 ;
Insert Into Options Set OptionId = 97 ,OptionTitle=  '850hPa' ,QuestionId=  25 ;
Insert Into Options Set OptionId = 98 ,OptionTitle=  '700hPa' ,QuestionId=  25 ;
Insert Into Options Set OptionId = 99 ,OptionTitle=  '500hPa' ,QuestionId=  25 ;
Insert Into Options Set OptionId = 100 ,OptionTitle=  '300hPa' ,QuestionId=  25 ;
Insert Into Options Set OptionId = 101 ,OptionTitle=  '对流层、等温层、中间层、热层、散逸层' ,QuestionId=  26 ;
Insert Into Options Set OptionId = 102 ,OptionTitle=  '对流层、平流层、中间层、热层、散逸层' ,QuestionId=  26 ;
Insert Into Options Set OptionId = 103 ,OptionTitle=  '对流层、平流层、中间层、散逸层、热层' ,QuestionId=  26 ;
Insert Into Options Set OptionId = 104 ,OptionTitle=  '散逸层、热层、中间层、平流层、对流层' ,QuestionId=  26 ;
