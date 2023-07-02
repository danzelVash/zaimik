import { TypeDayVariants } from './types';

export const getDayWord = (day: number): TypeDayVariants => {
	const lastDigit: number = day % 10;
	const lastTwoDigits: number = day % 100;

	if (lastTwoDigits >= 11 && lastTwoDigits <= 19) {
		return 'дней';
	}

	if (lastDigit % 10 === 1) {
		return 'день';
	}

	if (lastDigit >= 2 && lastDigit <= 4) {
		return 'дня';
	}

	return 'дней';
};

export const handleScrollBarPadding = (
	offset: number,
	...items: HTMLElement[]
): void => {
	items.forEach(item => {
		item.style.paddingRight = `${offset}px`;
	});
};
