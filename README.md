# Golang client SDK fo MoMo APIs

This Golang SDK client interfaces with MTN's MoMo API (Collections, Disbursements & Remittances ) to support building of great products on top of MTN MoMo payments around eight countries in Africa, including;

1. Uganda
2. Ghana
3. Ivory Coast
4. Zambia
5. Cameroon
6. Benin
7. Congo
8. Swaziland

## 1. Collections API

This API enables remote collection of bills, fees or taxes. Collections is a service that enables Mobile Money partners to receive payments for goods and services using MTN Mobile Money.

The services can be face-to-face like MomoPay or can be done remotely for both offline and online.
Payments can be customer-initiated on USSD/App/Web or Merchant-initiated where a customer is sent a debit request for approval.

Once the service is enabled, a Collections account is created for the partner to which funds are deposited/received.
The partner is able to make onward payments to their suppliers/partners/employees (B2B or B2C) and/or liquidate the collected
funds into their respective bank accounts.

The service offers the convenience of collection for online payments, bills, loan repayments, contribution to activities and installments to mutually agreed upon repayments of services and products. Examples of partners with bill collections are Umeme, National Water, URA, NSSF, DSTV to name but a few.

## 2. Disbursements API

This API automatically deposits funds to multiple users. Disbursements is a service that enables Mobile Money partners to send money in bulk to different recipients with just one click.

This setup can be manually executed (logging into the system, uploading recipient's list and trigger payments) or automated (requires a one-time setup of the recipients' lists and commands to effect payment).

Examples of partners that use this service are: Betting companies to pay winners, Disbursements of funds to refugees/beneficiaries among others.

The expectation is that the partner opens a Disbursement account with MTN and this account is pre-funded to enable the payments once the requests come through from the partner..

## 3. Remittances API

This API remits funds to local recipients from the diaspora with ease. Remittances is a solution that enables a customer to transfer or receive funds from the diaspora to Mobile Money recipient's account in local currency.

This is an automated solution where the money is transferred in real-time when the request hits the system (works in a similar way as the automated Disbursements solution).

The sender logs on the Web/App/USSD (USSD only for outbound) or visit the remitter’s point of sale to send money.

A request is then sent to the local partner’s system that triggers a payment into the recipient’s wallet.

The partner is required to open an account with MTN and load funds on it to facilitate the transfer of funds (inbound and outbound).

## Notice

This library is still under massive development and rapid prototyping.
