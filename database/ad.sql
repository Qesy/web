-- phpMyAdmin SQL Dump
-- version 3.5.6
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2013 年 03 月 22 日 10:56
-- 服务器版本: 5.5.29-0ubuntu0.12.04.1
-- PHP 版本: 5.3.10-1ubuntu3.5

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `ad`
--

-- --------------------------------------------------------

--
-- 表的结构 `ad_activity`
--

CREATE TABLE IF NOT EXISTS `ad_activity` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `activity_id` int(11) NOT NULL DEFAULT '0',
  `goods_id` int(11) NOT NULL DEFAULT '0',
  `ip` int(11) NOT NULL DEFAULT '0',
  `time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `ad_admin`
--

CREATE TABLE IF NOT EXISTS `ad_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '',
  `password` varchar(50) NOT NULL DEFAULT '',
  `permissions` tinyint(3) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `ad_admin`
--

INSERT INTO `ad_admin` (`id`, `username`, `password`, `permissions`) VALUES
(1, 'admin', 'admin', 1);

-- --------------------------------------------------------

--
-- 表的结构 `ad_email`
--

CREATE TABLE IF NOT EXISTS `ad_email` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `mail_id` int(11) NOT NULL DEFAULT '0',
  `goods_id` int(11) NOT NULL DEFAULT '0',
  `ip` int(11) NOT NULL DEFAULT '0',
  `time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `ad_ip`
--

CREATE TABLE IF NOT EXISTS `ad_ip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `adid` int(11) NOT NULL DEFAULT '0',
  `ip` bigint(20) NOT NULL DEFAULT '0',
  `time` int(11) NOT NULL DEFAULT '0',
  `type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '0:pic,1:mail,2:activity',
  `domain` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `refer` varchar(255) DEFAULT NULL,
  `goods_id` int(11) NOT NULL DEFAULT '0',
  `pv` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=17 ;

--
-- 转存表中的数据 `ad_ip`
--

INSERT INTO `ad_ip` (`id`, `adid`, `ip`, `time`, `type`, `domain`, `ua`, `refer`, `goods_id`, `pv`) VALUES
(1, 1, 3232239628, 1363046400, 0, '9yuejiu.test.com', 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.152 Safari/537.22', 'http://9yuejiu.test.com/category-3-b0.html', 0, 33),
(2, 0, 3232239633, 1363046400, 0, '', 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)', '', 0, 26),
(3, 0, 3232239619, 1363046400, 0, '', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.99 Safari/537.22', '', 0, 5),
(4, 0, 3232239628, 1363132800, 0, '9yuejiu.test.com', 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.152 Safari/537.22', 'http://9yuejiu.test.com/category-4-b0.html', 0, 142),
(5, 332, 2130706433, 1363132800, 0, '9yuejiu.jiangchuan.com', 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)', 'http://9yuejiu.jiangchuan.com/user.php?act=register', 0, 36),
(6, 0, 3232239619, 1363132800, 0, '', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.99 Safari/537.22', '', 0, 13),
(7, 0, 2130706433, 1363219200, 0, '', 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)', '', 0, 38),
(8, 332, 2130706433, 1363219200, 0, '', 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)', '', 0, 47),
(9, 0, 3232239619, 1363219200, 0, '', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.99 Safari/537.22', '', 0, 4),
(10, 0, 3232239628, 1363219200, 0, '', 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.152 Safari/537.22', '', 0, 44),
(11, 0, 3232239628, 1363305600, 0, '', 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.152 Safari/537.22', '', 0, 115),
(12, 0, 2130706433, 1363305600, 0, '', 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.152 Safari/537.22', '', 0, 191),
(13, 332, 2130706433, 1363305600, 0, '', 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)', '', 0, 12),
(14, 0, 3232239619, 1363305600, 0, '', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.99 Safari/537.22', '', 0, 3),
(15, 0, 3232239620, 1363305600, 0, '', 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.4 (KHTML, like Gecko) Chrome/22.0.1229.95 Safari/537.4', '', 0, 3),
(16, 0, 3232239633, 1363305600, 0, '', 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)', '', 0, 9);

-- --------------------------------------------------------

--
-- 表的结构 `ad_list`
--

CREATE TABLE IF NOT EXISTS `ad_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `img` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) DEFAULT NULL,
  `width` int(11) NOT NULL DEFAULT '0',
  `height` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(3) NOT NULL DEFAULT '1',
  `addtime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `ad_list`
--

INSERT INTO `ad_list` (`id`, `name`, `img`, `url`, `width`, `height`, `status`, `addtime`) VALUES
(1, '撒的发生', '345435', '345345', 34, 45, 1, 1363327479);

-- --------------------------------------------------------

--
-- 表的结构 `ad_login`
--

CREATE TABLE IF NOT EXISTS `ad_login` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `adid` int(11) NOT NULL,
  `uid` int(11) NOT NULL,
  `type` int(11) NOT NULL,
  `time` int(11) NOT NULL,
  `ip` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=10 ;

--
-- 转存表中的数据 `ad_login`
--

INSERT INTO `ad_login` (`id`, `adid`, `uid`, `type`, `time`, `ip`) VALUES
(1, 0, 94, 0, 1363219200, 2130706433),
(2, 0, 94, 0, 1363219200, 2130706433),
(3, 0, 94, 0, 1363305600, 2130706433),
(4, 0, 94, 0, 1363305600, 2130706433),
(5, 0, 94, 0, 1363305600, 2130706433),
(6, 0, 94, 0, 1363305600, 2130706433),
(7, 0, 94, 0, 1363305600, 3232239628),
(8, 0, 94, 0, 1363305600, 3232239628),
(9, 0, 94, 0, 1363305600, 3232239628);

-- --------------------------------------------------------

--
-- 表的结构 `ad_order`
--

CREATE TABLE IF NOT EXISTS `ad_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `adid` int(11) NOT NULL DEFAULT '0',
  `uid` int(11) NOT NULL DEFAULT '0',
  `order_id` bigint(20) NOT NULL DEFAULT '0',
  `money` decimal(10,2) NOT NULL DEFAULT '0.00',
  `real_money` decimal(10,2) NOT NULL DEFAULT '0.00',
  `time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `ad_order`
--

INSERT INTO `ad_order` (`id`, `adid`, `uid`, `order_id`, `money`, `real_money`, `time`) VALUES
(1, 1, 1, 234234, 3.00, 4.00, 1360857633);

-- --------------------------------------------------------

--
-- 表的结构 `ad_reg`
--

CREATE TABLE IF NOT EXISTS `ad_reg` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `adid` int(11) NOT NULL DEFAULT '0',
  `ip` bigint(20) NOT NULL DEFAULT '0',
  `uid` int(11) NOT NULL DEFAULT '0',
  `time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `ad_reg`
--

INSERT INTO `ad_reg` (`id`, `adid`, `ip`, `uid`, `time`) VALUES
(1, 332, 2130706433, 98, 1363132800);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
