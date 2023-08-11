export default function isObjectFalse(obj:{}) {
    if (!obj) return false;

		return Object.keys(obj).some((k) => {
			const value = obj[k as keyof typeof obj];
			if (typeof value === 'object') {
				return Object.keys(value).length > 0;
			} else return Boolean(value);
		});
}