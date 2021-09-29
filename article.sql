-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 29 Sep 2021 pada 11.52
-- Versi server: 10.1.38-MariaDB
-- Versi PHP: 7.3.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `article`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `articles`
--

CREATE TABLE `articles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext,
  `slug` longtext,
  `content` text,
  `category_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `articles`
--

INSERT INTO `articles` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `slug`, `content`, `category_id`) VALUES
(1, '2021-09-29 14:31:56.000', '2021-09-29 14:31:56.000', NULL, 'Perilisan Iphone 13', 'perilisan-iphone-13', 'Iphone 13 akan dirilis pada tahun 2021', 2),
(2, '2021-09-29 14:42:04.000', '2021-09-29 15:55:16.304', NULL, 'Iphone 12 rilis Rev 1.1', 'iphone-12-rilis', 'perilisan iphone 12 di tahun xxx', 4),
(3, '2021-09-29 15:15:41.297', '2021-09-29 15:15:41.297', '2021-09-29 15:56:03.626', 'test', 'test', 'content tes', 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `categories`
--

CREATE TABLE `categories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `category_name` longtext,
  `category_slug` longtext
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `categories`
--

INSERT INTO `categories` (`id`, `created_at`, `updated_at`, `deleted_at`, `category_name`, `category_slug`) VALUES
(1, '2021-09-29 13:34:04.485', '2021-09-29 13:34:33.175', NULL, 'android new', 'android'),
(2, '2021-09-29 13:34:18.453', '2021-09-29 13:34:18.453', NULL, 'iphone', 'iphone'),
(3, '2021-09-29 13:37:38.125', '2021-09-29 15:14:13.191', NULL, 'test', 'test'),
(4, '2021-09-29 13:51:51.802', '2021-09-29 13:51:51.802', NULL, 'iphone 12', 'iphone12'),
(5, '2021-09-29 15:10:59.668', '2021-09-29 15:10:59.668', NULL, 'iphone 13', 'iphone13'),
(6, '2021-09-29 15:12:13.294', '2021-09-29 15:12:13.294', NULL, 'iphone 14', 'iphone14'),
(7, '2021-09-29 15:13:47.105', '2021-09-29 15:13:47.105', NULL, 'iphone 15', 'iphone15');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_articles_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_categories_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `articles`
--
ALTER TABLE `articles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `categories`
--
ALTER TABLE `categories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
