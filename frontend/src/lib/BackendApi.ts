import Config from "@/Config";

export async function backendApiGet(
  actionIdentifier: string,
  callbacks?: {
    onLoad?: (l: boolean) => void;
    onError?: (e?: Error) => void;
    onSuccess?: (returnValue: object) => void;
  },
): Promise<object | null> {
  const url = Config.Api.Combine(actionIdentifier);

  callbacks?.onLoad?.(true);
  callbacks?.onError?.();

  try {
    const response = await fetch(url, {
      //credentials: 'include',
    });

    if (!response.ok) {
      throw new Error((await response.text()).substring(0, 500));
    }

    const returnValue = await response.json();
    callbacks?.onSuccess?.(returnValue);
    return returnValue;
  } catch (e) {
    callbacks?.onError?.(e as Error);
  } finally {
    callbacks?.onLoad?.(false);
  }

  return null;
}

export async function backendApiPost(
  actionIdentifier: string,
  callbacks: {
    onLoad?: (l: boolean) => void;
    onError?: (e?: Error) => void;
    onSuccess?: (returnValue: any) => void;
  },
  payload?: object,
): Promise<object | null> {
  const url = Config.Api.Combine(actionIdentifier);

  callbacks?.onLoad?.(true);
  callbacks?.onError?.();

  try {
    const response = await fetch(url, {
      //credentials: 'include',
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      throw new Error((await response.text()).substring(0, 500));
    }

    const returnValue = await response.json();
    callbacks?.onSuccess?.(returnValue);
    return returnValue;
  } catch (e) {
    callbacks?.onError?.(e as Error);
  } finally {
    callbacks?.onLoad?.(false);
  }

  return null;
}
