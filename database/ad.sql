-- phpMyAdmin SQL Dump
-- version 3.3.7
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2013 年 02 月 16 日 09:35
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
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- 转存表中的数据 `admin`
--


-- --------------------------------------------------------

--
-- 表的结构 `ad_activity`
--

CREATE TABLE IF NOT EXISTS `ad_activity` (
  `id` int(11) NOT NULL auto_increment,
  `uid` int(11) NOT NULL default '0',
  `activity_id` int(11) NOT NULL default '0',
  `goods_id` int(11) NOT NULL default '0',
  `ip` int(11) NOT NULL default '0',
  `time` int(11) NOT NULL default '0',
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED AUTO_INCREMENT=1 ;

--
-- 转存表中的数据 `ad_activity`
--


-- --------------------------------------------------------

--
-- 表的结构 `ad_email`
--

CREATE TABLE IF NOT EXISTS `ad_email` (
  `id` int(11) NOT NULL auto_increment,
  `uid` int(11) NOT NULL default '0',
  `mail_id` int(11) NOT NULL default '0',
  `goods_id` int(11) NOT NULL default '0',
  `ip` int(11) NOT NULL default '0',
  `time` int(11) NOT NULL default '0',
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- 转存表中的数据 `ad_email`
--


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
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- 转存表中的数据 `ad_ip`
--


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
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- 转存表中的数据 `ad_list`
--


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
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- 转存表中的数据 `ad_order`
--


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
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- 转存表中的数据 `ad_reg`
--

