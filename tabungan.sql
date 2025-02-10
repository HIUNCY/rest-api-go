-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Feb 10, 2025 at 05:45 PM
-- Server version: 8.0.30
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tabungan`
--

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id_tabungan` bigint NOT NULL,
  `nik` varchar(25) NOT NULL,
  `setor` bigint NOT NULL,
  `tarik` bigint NOT NULL,
  `tanggal` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `jenis` enum('setor','tarik') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id_tabungan`, `nik`, `setor`, `tarik`, `tanggal`, `jenis`) VALUES
(1, '3172031313131', 10000, 0, '2025-02-10 17:33:40', 'setor'),
(2, '3172031313131', 0, 10000, '2025-02-10 17:37:46', 'tarik');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `nik` varchar(25) NOT NULL,
  `username` varchar(191) NOT NULL,
  `password` varchar(100) NOT NULL,
  `full_name` varchar(100) NOT NULL,
  `email` varchar(191) NOT NULL,
  `phone_number` longtext,
  `role` enum('admin','nasabah') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`nik`, `username`, `password`, `full_name`, `email`, `phone_number`, `role`) VALUES
('3172031313131', 'ningga', '$2a$10$8jeiJ/bOtl7QhLK/uoZGO.EVGqAxWNC/piNfX07LtQyjt6vabxNBe', 'rahmat rider', 'rahmat@gmail.com', '0896886886', 'nasabah');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id_tabungan`),
  ADD KEY `nik` (`nik`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`nik`),
  ADD UNIQUE KEY `uni_users_email` (`email`),
  ADD UNIQUE KEY `uni_users_username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id_tabungan` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
