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
(3, 18, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 08:31:15', '2020-10-08 08:31:15'),
(7, 35925, 'UNASSIGNED', '-6.362764', '106.827049', '-6.238270', '106.975571', '2020-10-08 11:31:51', '2020-10-08 11:31:51');

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
