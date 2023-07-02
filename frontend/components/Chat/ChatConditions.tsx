const ChatConditions = () => {
	return (
		<div
			className='md:mt-6 mt-5 grid lg:grid-cols-2 grid-cols-1 md:gap-3 text-center text-white bg-tertiary lg:p-5 p-3 rounded-3xl shadow-[0px_5px_10px_rgba(0,0,0,.25)] font-bold leading-tight'
		>
			<div className='grid md:grid-cols-2 grid-cols-1 gap-3 md:py-0 py-3 bg-accent rounded-3xl shadow-[0px_5px_10px_rgba(0,0,0,.25)]'>
				<div className='lg:px-5 md:px-3 md:py-3 lg:space-y-4 md:space-y-3 space-y-1.5'>
					<div className='xl:text-[24px] md:text-[21px] text-[20px]'>Сумма займа</div>
					<div className='uppercase xl:text-[36px] lg:text-[26px] md:text-[23px] text-[20px]'>9999 руб</div>
				</div>
				<div className='lg:px-5 md:px-3 md:py-3 lg:space-y-4 md:space-y-3 space-y-1.5'>
					<div className='xl:text-[24px] md:text-[21px] text-[20px]'>Срок</div>
					<div className='uppercase xl:text-[36px] lg:text-[26px] md:text-[23px] text-[20px]'>1 день</div>
				</div>
			</div>
			<div className='grid md:grid-cols-2 grid-cols-1 gap-3 md:py-0 py-3'>
			<div className='lg:px-5 md:px-3 md:py-3 lg:space-y-4 md:space-y-3 space-y-1.5'>
					<div className='xl:text-[24px] md:text-[21px] text-[20px]'>Вы возвращаете</div>
					<div className='uppercase xl:text-[36px] lg:text-[26px] md:text-[23px] text-[20px]'>10 100 РУБ</div>
				</div>
				<div className='lg:px-5 md:px-3 md:py-3 lg:space-y-4 md:space-y-3 space-y-1.5'>
					<div className='xl:text-[24px] md:text-[21px] text-[20px]'>Дата возврата</div>
					<div className='uppercase xl:text-[36px] lg:text-[26px] md:text-[23px] text-[20px]'>01.01.2024</div>
				</div>
			</div>
		</div>
	)
}

export default ChatConditions;