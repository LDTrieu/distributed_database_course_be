
/****** Object:  StoredProcedure [dbo].[SP_GET_CauHoi]    Script Date: 4/24/2022 10:06:17 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[SP_GET_CauHoi]
	@maMH nchar(5), @trinhDo nchar(1), @soCauThi int
AS
BEGIN
	DECLARE @countCauHoi int, @countCauHoiSiteKhac int, @TrinhDoDuoi nchar(1),
	 @countCauHoiDuoi int, @countCHDuoiSiteKhac int
	
	--Trình độ A
	IF(@trinhDo = 'A')
	BEGIN 
		SET @TrinhDoDuoi = 'B'
	END
	--Trình độ B
	ELSE IF(@trinhDo = 'B')
	BEGIN 
		SET @TrinhDoDuoi = 'C'
	END
	IF(@trinhDo = 'C')
	BEGIN
		-- Đếm số câu hỏi có môn học, trình độ là các tham số và thuộc giáo viên của site hiện tại
		SELECT @countCauHoi = COUNT(CAUHOI) FROM BODE 
		WHERE MAMH = 'AVCB' AND TRINHDO = 'A' AND MAGV IN 
		(SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA)) 
			
		-- Nếu số câu hỏi đó lớn hơn hoặc bằng số câu cần lấy
		IF(@countCauHoi >= @soCauThi)
		BEGIN
			-- Lấy ngẫu nhiên số câu hỏi trong bộ đề thuộc site hiện tại, , NewID() is used to generate unique identifier value
			SELECT TOP (@soCauThi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
			WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
			ORDER BY  NEWID()
		END
		-- Nếu số câu hỏi đó bé hơn số câu cần lấy
		ELSE IF(@countCauHoi < @soCauThi)
		BEGIN
			-- Đếm số câu hỏi thuộc site khác
			SELECT @countCauHoiSiteKhac = COUNT(CAUHOI) FROM BODE 
				WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN 
				(SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA)) 

			IF(@countCauHoiSiteKhac < @soCauThi - @countCauHoi)
			BEGIN-- thiếu đề
				RAISERROR('Không đủ số câu để tạo đề!', 16, 1)
			END
			ELSE IF(@countCauHoiSiteKhac >= @soCauThi - @countCauHoi)
			BEGIN
				--UNION 2 SELECT random câu hỏi của giáo viên cả 2 site
				SELECT * FROM ( 
					SELECT TOP (@countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
					WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
					ORDER BY NEWID()
				) AS SITEHT
				UNION ALL
				SELECT * FROM ( 
					SELECT TOP (@soCauThi - @countCauHoiSiteKhac) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
					WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
					ORDER BY  NEWID()
				) AS SITEKHAC
			END
		END
	END
	-- Trình độ A hoặc B
	ELSE 
	BEGIN
		--Đếm số câu hỏi site hiện tại (A)
		SELECT @countCauHoi = COUNT(CAUHOI) FROM BODE 
		WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
		
		IF(@countCauHoi >= @soCauThi)--Nếu đủ (A>= 100%)
		BEGIN
			SELECT TOP (@soCauThi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
			WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
			ORDER BY  NEWID()
		END
		ELSE IF(@countCauHoi >= @soCauThi*0.7) -- Trên 70% số câu hỏi cần thì lấy ở trình độ dưới
		BEGIN 
			SELECT @countCauHoiDuoi = COUNT(CAUHOI) FROM BODE -- Đếm số lượng câu hỏi ở trình độ dưới (B)
			WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
					
			-- Nếu câu thi ở trình độ dưới đủ (B >= 100% - A)
			IF(@countCauHoiDuoi >= @soCauThi - @countCauHoi) 
				BEGIN 
					SELECT * FROM (
						SELECT TOP (@countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TRINHDOTREN
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TRINHDODUOI
				END
			ELSE -- Nếu câu thi ở trình độ dưới không đủ qua cơ sở khác
			BEGIN
				-- Đếm câu hỏi cùng trình độ mà ở site khác (A2)
				SELECT @countCauHoiSiteKhac = COUNT(CAUHOI) FROM BODE 
				WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						
				-- Đếm câu hỏi trình độ dưới ở site khác (B2)
				SELECT @countCHDuoiSiteKhac = COUNT(CAUHOI) FROM BODE
				WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
							
				-- Nếu câu hỏi cùng trình độ site khác đủ ( A2 >=  100% - A - B)
				IF(@countCauHoiSiteKhac >= @soCauThi - @countCauHoi - @countCauHoiDuoi )
				BEGIN
					SELECT * FROM (
						SELECT TOP (@countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TRINHDOTREN
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TRINHDODUOI
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi - @countCauHoiDuoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TRINHDOSITEKHAC
				END
				-- Nếu mà câu hỏi cùng trình độ site khác mà không đủ thì lấy câu hỏi trình độ dưới site khác(B2 >= 100 -A - A2 - B)
				ELSE IF(@countCauHoiSiteKhac >= @soCauThi - @countCauHoi - @countCauHoiSiteKhac - @countCauHoiDuoi)
				BEGIN
					SELECT * FROM (
						SELECT TOP (@countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDoTren
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDoDuoi
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi - @countCauHoiDuoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDoSiteKhac
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi - @countCauHoiDuoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDoDuoiSiteKhac
				END 
				-- Nếu không đủ
				ELSE 
				BEGIN-- thiếu đề
					RAISERROR('Không đủ số câu để tạo đề!', 16, 1)
				END
			END
		END
		-- Nếu mà số câu hỏi cùng trình độ mà không đủ 70% thì qua site khác lấy
		ELSE 
		BEGIN 
			-- Đếm câu hỏi cùng trình độ mà ở site khác (A2)
			SELECT @countCauHoiSiteKhac = COUNT(CAUHOI) FROM BODE 
			WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
			
			-- Nếu câu hỏi site khác cùng trình độ đủ
			IF(@countCauHoiSiteKhac >= @soCauThi - @countCauHoi)
			BEGIN 
				SELECT * FROM (
					SELECT TOP (@countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
					WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
					ORDER BY  NEWID()
				) AS CauHoiCungSite
				UNION ALL
				SELECT * FROM (
					SELECT TOP (@soCauThi - @countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
					WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
					ORDER BY  NEWID()
				) AS CungTrinhDoSiteKhac
			END	
			-- Nếu mà câu hỏi cùng trình độ site khác và câu hỏi cùng trình độ site hiện tại 
			-- tổng lớn hơn 70% thì 30% còn lại lấy câu hỏi trình độ dưới (A2 + A >= 70%)
			ELSE IF(@countCauHoiSiteKhac >= @soCauThi*0.7 - @countCauHoi) 
			BEGIN
				SELECT @countCauHoiDuoi = COUNT(CAUHOI) FROM BODE -- B
				WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						
				SELECT @countCHDuoiSiteKhac = COUNT(CAUHOI) FROM BODE -- B2
				WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
					
				-- Lấy câu hỏi trình độ dưới cùng site 
				IF(@countCauHoiDuoi>= @soCauThi - @countCauHoi - @countCauHoiSiteKhac ) 
					BEGIN
						SELECT * FROM (
						SELECT TOP (@countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDo
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS CungTrinhDoSiteKhac
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi - @countCauHoiDuoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDoDuoi
				END
				-- Lấy câu hỏi trình độ dưới khác site thêm vào( B2 >= 100 -A - A2 - B)
				ELSE IF(@countCHDuoiSiteKhac >= @soCauThi - @countCauHoi - @countCauHoiSiteKhac - @countCauHoiDuoi)
				BEGIN 
					SELECT * FROM (
						SELECT TOP (@countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDo
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @trinhDo AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS CungTrinhDoSiteKhac
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi - @countCauHoiDuoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDoDuoi
					UNION ALL
					SELECT * FROM (
						SELECT TOP (@soCauThi - @countCauHoi - @countCauHoiDuoi) CAUHOI, NOIDUNG, A,B,C,D,DAP_AN FROM BODE 
						WHERE MAMH = @maMH AND TRINHDO = @TrinhDoDuoi AND MAGV NOT IN (SELECT MAGV FROM GIAOVIEN WHERE MAKH IN (SELECT MAKH FROM KHOA))  
						ORDER BY  NEWID()
					) AS TrinhDoDuoiSiteKhac
				END
				ELSE -- thiếu đề
				BEGIN
					RAISERROR('Không đủ số câu để tạo đề!', 16, 1)
				END
			END
			ELSE -- thiếu đề
			BEGIN
				RAISERROR('Không đủ số câu để tạo đề!', 16, 1)
			END
				
		END
	END

END
GO