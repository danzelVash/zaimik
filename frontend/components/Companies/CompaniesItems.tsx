import Image from 'next/image';
import Link from 'next/link';

const CompaniesItems: React.FC = () => {
	return (
		<div className='scrollbar scrollbar-thumb-accent scrollbar-track-[#D9D9D9] xl:max-h-[977px] lg:max-h-[873px] md:max-h-[1604px] max-h-[1088px] overflow-auto mt-6 bg-tertiary md:py-7 py-5 2xl:px-9 md:px-7 px-3 rounded-[30px] grid xl:grid-cols-3 lg:grid-cols-2 grid-cols-1 2xl:gap-x-10 gap-x-7 gap-y-7'>
			{Array(12)
				.fill('')
				.map((_, index) => (
					<div
						key={index}
						className='rounded-[30px] overflow-hidden shadow-[0px_5px_10px_rgba(0,0,0,.25)]'
					>
						<div>
							<Image
								className='xl:h-[153px] lg:h-[133px] md:h-[190px] h-[130px] w-full object-cover'
								src='/static/search/01.jpg'
								width={426}
								height={153}
								alt=''
							/>
						</div>
						<div className='main-gradient lg:px-5 md:px-8 px-5 xl:py-7 lg:py-5 md:py-7 py-3.5 flex items-center justify-between gap-4 text-white leading-tight'>
							<div>
								<div className='lg:text-[24px] md:text-[30px] text-[16px] font-bold'>КредитЗайм</div>
								<div className='mt-1 lg:text-[18px] md:text-[26px] text-[16px] font-extrabold'>
									До 25 000 рублей
								</div>
								<div className='mt-1 lg:text-[16px] md:text-[22px] text-[12px]'>от 0% в день</div>
							</div>
							<Link
								className='flex items-center justify-center text-center shadow-[0px_5px_10px_rgba(0,0,0,.25)] text-white lg:text-[22px] md:text-[28px] text-[16px] font-extrabold lg:h-[63px] md:h-[120px] h-[58px] rounded-[26px] bg-accent w-full lg:max-w-[155px] md:max-w-[230px] max-w-[133px] transition-colors hover:bg-[#dc7c2a] duration-300'
								href=''
							>
								Получить
							</Link>
						</div>
					</div>
				))}
		</div>
	);
};

export default CompaniesItems;
