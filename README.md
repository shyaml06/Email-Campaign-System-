# 📧 Email Campaign System

A lightweight, high-performance **Email Campaign System** written in **Go (Golang)** for managing and sending bulk or targeted email campaigns efficiently.  
This system uses a **producer–consumer architecture** and supports SMTP integration, making it ideal for small to medium-scale email automation or outreach projects.

---

## 🚀 Features

- ✅ Queue-based producer–consumer system for sending emails concurrently  
- ✅ SMTP-based email sending with configurable credentials  
- ✅ JSON campaign input for flexible campaign setup  
- ✅ Simple, modular Go codebase (easy to extend)  
- ✅ Logs and error-handling for tracking campaign progress  

---

## 🏗️ Project Structure

Email-Campaign-System/
│
├── main.go # Entry point – initializes producers & consumers
├── producer.go # Reads email data and sends tasks to queue
├── consumer.go # Consumes queued tasks and sends emails
├── smtp.go # SMTP connection, authentication & email sending
├── campaign.json # Sample campaign input file
└── README.md # Project documentation


🧱 Architecture Overview

+---------------------+
|   campaign.json     |
+----------+----------+
           |
           v
+---------------------+
|     Producer        |
|  (reads & queues)   |
+----------+----------+
           |
           v
+---------------------+
|     Consumer(s)     |
| (send emails via    |
|   SMTP concurrently)|
+----------+----------+
           |
           v
+---------------------+
|    SMTP Server      |
+---------------------+
