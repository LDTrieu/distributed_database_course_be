USE TN_CSDLPT
go
CREATE VIEW V_DS_PHANMANH
AS
	SELECT TENCN=PUBS.description, TENSERVER=subscriber_server
	FROM sysmergepublications PUBS,sysmergesubscriptions SUBS
	WHERE PUBS.pubid =SUBS.pubid AND publisher <> subscriber_server
GO
SELECT * from  V_DS_PHANMANH 