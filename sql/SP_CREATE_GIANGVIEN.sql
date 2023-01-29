/****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP (1000) [MAGV]
      ,[HO]
      ,[TEN]
      ,[DIACHI]
      ,[MAKH]
      ,[rowguid]
  FROM [TN_CSDLPT].[dbo].[GIAOVIEN]

USE TN_CSDLPT
GO

CREATE PROCEDURE [dbo].[SP_CREATE_GIANGVIEN]
@MAGV  CHAR(8),
@HO  NVARCHAR(50),
@TEN  NVARCHAR(10),
@DIACHI  NVARCHAR(50),
@MAKH  NCHAR(8)

AS
DECLARE @RET INT
BEGIN
	INSERT INTO GIAOVIEN (MAGV,HO,TEN,DIACHI,MAKH) VALUES (@MAGV,@HO,@TEN,@DIACHI,@MAKH)

	END
	GO

DECLARE @result int
 @result = EXEC [SP_CREATE_GIANGVIEN] 