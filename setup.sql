create table users
(
  id           varchar(100),
  display_name varchar(25),

  primary key (id)
);

create table scopes
(
  id    int(11) auto_increment,
  scope varchar(50) unique,

  primary key (id)
);

create table user_tokens
(
  id            bigint auto_increment,
  user_id       varchar(100),
  access_token  varchar(30),
  refresh_token varchar(50),
  expires_in    int(11),

  primary key (id),
  foreign key (user_id) references users (id)
);

create table user_token_scope
(
  user_token_id bigint,
  scope_id      int,

  primary key (user_token_id, scope_id),
  foreign key (user_token_id) references user_tokens (id),
  foreign key (scope_id) references scopes (id)
);

create table notification_type
(
  id           int auto_increment,
  notification varchar(255),
  primary key (id)
);

create table notifications
(
  id                   bigint       not null auto_increment,

  user_id              varchar(100) not null,
  notification_type_id int          not null,

  notified             bit(1)       not null default false,

  created_at           datetime     not null default current_timestamp,

  primary key (id),
  foreign key (user_id) references users (id),
  foreign key (notification_type_id) references notification_type (id)
);

create table notification_subs
(
  notifications_id bigint,

  user_id          varchar(100) not null,
  username         varchar(25)  not null,
  display_name     varchar(25)  not null,

  type             varchar(10)  not null,
  sub_plan         varchar(50)  not null,
  months           smallint     not null,

  primary key (notifications_id),
  foreign key (notifications_id) references notifications (id)
);

create table notification_gift_subs
(
  notifications_id    bigint,

  user_id             varchar(100)         not null,
  username            varchar(25)          not null,
  display_name        varchar(25)          not null,

  anonymous           bit(1) default false not null,

  sub_plan            varchar(50)          not null,
  months              smallint             not null,

  gifter_id           varchar(100),
  gifter_username     varchar(25),
  gifter_display_name varchar(25),

  primary key (notifications_id),
  foreign key (notifications_id) references notifications (id)
);

create table notification_bits
(
  notifications_id  bigint,

  user_id           varchar(100) not null,
  username          varchar(25)  not null,
  display_name      varchar(25)  not null,

  bit_used          int          not null,
  total_bits_used   int          not null,

  message           varchar(500) not null,
  message_id        varchar(50)  not null,

  badge_entitlement int          not null,

  primary key (notifications_id),
  foreign key (notifications_id) references notifications (id)
);

create table notification_followers
(
  notifications_id bigint,

  user_id          varchar(100) not null,
  username         varchar(25)  not null,
  display_name     varchar(25)  not null,

  primary key (notifications_id),
  foreign key (notifications_id) references notifications (id)
);