# 📌 抖音商城（Tiktok Mall）

## 项目背景

随着移动互联网的普及和消费者购物习惯的变化，社交电商呈现出蓬勃发展的趋势。抖音作为一款拥有庞大用户群体的短视频社交平台，具有巨大的电商潜力。通过搭建电商平台，抖音可以为用户提供更加丰富的购物体验，同时为商家提供新的销售渠道，实现用户、商家和平台的多赢局面。

## 🚀 项目简介

本项目是一个基于 **微服务架构** 的系统，使用 **Go** 进行开发。
涉及技术栈：**gRPC**, **MySQL**, **Redis**, **RabbitMQ**, **Elasticsearch**。
使用 **Docker** 进行容器化部署，使用 **Kubernetes** 进行集群管理。

## 📐 架构概览

系统由多个微服务组成：

- **Admin Service**（管理微服务）：负责商品管理等后台管理功能。
- **Gateway Service**（网关微服务）：负责接收用户请求，进行鉴权等操作，然后转发给其他微服务。
- **Order Service**（订单微服务）：负责处理用户下单等订单相关操作。
- **Payment Service**（支付微服务）：负责处理用户支付等支付相关操作。
- **User Service**（用户微服务）：负责用户购物车、浏览商品等用户相关操作。
