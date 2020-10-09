-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Oct 08, 2020 at 01:47 PM
-- Server version: 10.4.8-MariaDB
-- PHP Version: 7.3.10

use delivery_app;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `delivery_app`
--

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `distance` int(11) NOT NULL,
  `status` varchar(30) NOT NULL,
  `start_latitude` varchar(100) NOT NULL,
  `start_longitude` varchar(100) NOT NULL,
  `end_latitude` varchar(100) NOT NULL,
  `end_longitude` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `distance`, `status`, `start_latitude`, `start_longitude`, `end_latitude`, `end_longitude`, `created_at`, `updated_at`) VALUES
(1, 18, 'UNASSIGNED', 'ss', 'dd', 'sss', 'fg', '2020-10-08 08:08:10', '2020-10-08 08:08:10'),
(2, 18, 'UNASSIGNED', 'ss', 'dd', 'sss', 'fg', '2020-10-08 08:28:19', '2020-10-08 08:28:19'),
(3, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:31:15', '2020-10-08 08:31:15'),
(4, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:32:38', '2020-10-08 08:32:38'),
(5, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:33:20', '2020-10-08 08:33:20'),
(6, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:33:50', '2020-10-08 08:33:50'),
(7, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:36:25', '2020-10-08 08:36:25'),
(8, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:39:43', '2020-10-08 08:39:43'),
(9, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:44:08', '2020-10-08 08:44:08'),
(10, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:46:43', '2020-10-08 08:46:43'),
(11, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:47:13', '2020-10-08 08:47:13'),
(12, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:48:05', '2020-10-08 08:48:05'),
(13, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:49:37', '2020-10-08 08:49:37'),
(14, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:50:10', '2020-10-08 08:50:10'),
(15, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:52:59', '2020-10-08 08:52:59'),
(16, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:53:52', '2020-10-08 08:53:52'),
(17, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:54:07', '2020-10-08 08:54:07'),
(18, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:54:58', '2020-10-08 08:54:58'),
(19, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:55:43', '2020-10-08 08:55:43'),
(20, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:56:55', '2020-10-08 08:56:55'),
(21, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:57:43', '2020-10-08 08:57:43'),
(22, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:58:13', '2020-10-08 08:58:13'),
(23, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:59:26', '2020-10-08 08:59:26'),
(24, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:59:51', '2020-10-08 08:59:51'),
(25, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 09:00:16', '2020-10-08 09:00:16'),
(26, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 09:00:35', '2020-10-08 09:00:35'),
(27, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 09:00:49', '2020-10-08 09:00:49'),
(28, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 09:01:00', '2020-10-08 09:01:00'),
(29, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 09:26:32', '2020-10-08 09:26:32'),
(30, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 09:26:49', '2020-10-08 09:26:49'),
(31, 35925, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 11:31:51', '2020-10-08 11:31:51');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
