const ReviewCard = () => {
	return (
		<div className='main-gradient md:p-5 p-4 rounded-2xl'>
			<div className='lg:text-[27px] md:text-[24px] text-[19px] font-bold leading-tight flex flex-row flex-wrap gap-x-2 gap-y-1'>
				<span>Андрей,</span>
				<span>+7 (999) *** - ** - 33</span>
			</div>
			<p className='lg:mt-3 mt-2 font-medium lg:text-[24px] md:text-[21px] text-[16px] leading-tight'>
				Была необходимость взять займ в течение дня на 20 000 рублей, поэтому
				решил обратиться в Займик. Работа сервиса удивила своей скоростью. За 5
				минут я успел авторизоваться и получить подходящие варианты с самыми
				лучшими условиями. Однозначно рекомендую Займика!!
			</p>
		</div>
	);
};

export default ReviewCard;
