-- phpMyAdmin SQL Dump
-- version 3.3.7
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2013 年 01 月 25 日 05:27
-- 服务器版本: 5.0.90
-- PHP 版本: 5.2.14

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `ad`
--

-- --------------------------------------------------------

--
-- 表的结构 `admin`
--

CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(11) NOT NULL auto_increment,
  `username` varchar(50) NOT NULL default '',
  `password` varchar(50) NOT NULL default '',
  `permissions` tinyint(3) NOT NULL default '0',
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=3 ;

--
-- 转存表中的数据 `admin`
--

INSERT INTO `admin` (`id`, `username`, `password`, `permissions`) VALUES
(1, 'admin', 'admin', 1),
(2, 'admin', 'admin', 9);

-- --------------------------------------------------------

--
-- 表的结构 `ad_ip`
--

CREATE TABLE IF NOT EXISTS `ad_ip` (
  `id` int(11) NOT NULL auto_increment,
  `adid` int(11) NOT NULL default '0',
  `ip` int(11) NOT NULL default '0',
  `time` date NOT NULL default '0000-00-00',
  `type` tinyint(3) NOT NULL default '0' COMMENT '0:pic,1:mail,2:activity',
  `domain` varchar(50) default NULL,
  `ua` varchar(255) default NULL,
  `refer` varchar(255) default NULL,
  `goods_id` int(11) NOT NULL default '0',
  `pv` int(11) NOT NULL default '1',
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `ad_ip`
--

INSERT INTO `ad_ip` (`id`, `adid`, `ip`, `time`, `type`, `domain`, `ua`, `refer`, `goods_id`, `pv`) VALUES
(1, 1, -899624571, '2013-01-08', 0, 'www.q-cms.cn', 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11', '', 0, 945);

-- --------------------------------------------------------

--
-- 表的结构 `ad_list`
--

CREATE TABLE IF NOT EXISTS `ad_list` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(100) NOT NULL default '',
  `img` varchar(255) NOT NULL default '',
  `url` varchar(255) default NULL,
  `width` int(11) NOT NULL default '0',
  `height` int(11) NOT NULL default '0',
  `status` tinyint(3) NOT NULL default '1',
  `addtime` int(11) NOT NULL default '0',
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=10 ;

--
-- 转存表中的数据 `ad_list`
--

INSERT INTO `ad_list` (`id`, `name`, `img`, `url`, `width`, `height`, `status`, `addtime`) VALUES
(1, 'aaa', '/static/upload/1355057768.jpg', 'qqq', 200, 100, 1, 1355057768),
(2, 'aaa', '/static/upload/1355057872.jpg', 'ww', 200, 100, 1, 1355057872),
(3, 'aaa', '/static/upload/1355057918.jpg', 'eee', 200, 100, 1, 1355057918),
(4, 'aaa', '/static/upload/1355057967.jpg', 'rrr', 200, 100, 1, 1355057967),
(5, 'sss', '/static/upload/1355119255.jpg', 'tt', 200, 100, 1, 1355119255),
(6, '岁的范德萨', '/static/upload/1355131466.jpg', 'yy', 200, 100, 1, 1355131466),
(7, 'aaa', '/static/upload/1355885601.jpg', 'uuu', 441, 388, 1, 1355885601),
(8, 'asdf', '/static/upload/1355885641.jpg', 'asdfasdf', 23, 32, 1, 1355885641),
(9, 'Qesy百度', '/static/upload/1356930052.jpg', 'http://www.baidu.com', 400, 200, 1, 1356930097);

-- --------------------------------------------------------

--
-- 表的结构 `ad_order`
--

CREATE TABLE IF NOT EXISTS `ad_order` (
  `id` int(11) NOT NULL auto_increment,
  `adid` int(11) NOT NULL default '0',
  `uid` int(11) NOT NULL default '0',
  `order_id` bigint(20) NOT NULL default '0',
  `money` decimal(10,2) NOT NULL default '0.00',
  `real_money` decimal(10,2) NOT NULL default '0.00',
  `time` int(11) NOT NULL default '0',
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=13 ;

--
-- 转存表中的数据 `ad_order`
--

INSERT INTO `ad_order` (`id`, `adid`, `uid`, `order_id`, `money`, `real_money`, `time`) VALUES
(1, 1, 1234, 123456, 99.01, 80.99, 2013),
(2, 1, 1234, 123456, 99.01, 80.99, 1357636938),
(3, 1, 1234, 123456, 99.01, 80.99, 1357636938),
(4, 1, 1234, 123456, 99.01, 80.99, 1357636938),
(5, 1, 1234, 123456, 99.01, 80.99, 1357636938),
(6, 1, 1234, 123456, 99.01, 80.99, 1357636938),
(7, 1, 1234, 123456, 99.01, 80.99, 1357636938),
(8, 1, 1234, 123456, 99.01, 80.99, 1357636939),
(9, 1, 1234, 123456, 99.01, 80.99, 1357636939),
(10, 1, 1234, 123456, 99.01, 80.99, 1357636939),
(11, 1, 1234, 123456, 99.01, 80.99, 1357636939),
(12, 1, 1234, 123456, 99.01, 80.99, 1357636939);

-- --------------------------------------------------------

--
-- 表的结构 `ad_reg`
--

CREATE TABLE IF NOT EXISTS `ad_reg` (
  `id` int(11) NOT NULL auto_increment,
  `adid` int(11) NOT NULL default '0',
  `ip` int(11) NOT NULL default '0',
  `uid` int(11) NOT NULL default '0',
  `time` date NOT NULL default '0000-00-00',
  PRIMARY KEY  (`id`),
  UNIQUE KEY `uid` (`uid`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `ad_reg`
--

INSERT INTO `ad_reg` (`id`, `adid`, `ip`, `uid`, `time`) VALUES
(1, 1, -899624571, 1234, '2013-01-08');
