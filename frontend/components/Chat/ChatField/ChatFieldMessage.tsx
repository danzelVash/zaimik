import { ReactNode } from 'react'

interface IChatFieldMessageProps {
	children: ReactNode;
	position?: 'start' | 'end';
	className?: string;
}

const ChatFieldMessage: React.FC<IChatFieldMessageProps> = ({
	children,
	position = 'start',
	className = '',
}) => {
	const currentDate: Date = new Date();
	const hours: string = currentDate.getHours().toString().padStart(2, '0');
	const minutes: string = currentDate.getMinutes().toString().padStart(2, '0');

	return (
		<div className={`flex justify-${position}`}>
			<div
				className={`flex flex-col rounded-xl w-auto relative md:pl-4 pl-3 md:pt-4 pt-3 md:pr-10 pr-5 md:pb-7 pb-6 lg:max-w-[600px] md:max-w-[70%] max-w-[90%] font-medium md:text-[22px] text-[16px] leading-tight ${className}`}
			>
				<div className='inline-block'>{children}</div>
				<div className='absolute bottom-[8px] right-[8px] md:text-[16px] text-[14px] leading-none'>
					{hours}:{minutes}
				</div>
			</div>
		</div>
	);
};

export default ChatFieldMessage;
